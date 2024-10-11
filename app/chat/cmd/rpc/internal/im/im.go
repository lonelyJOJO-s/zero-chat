package im

import (
	"fmt"
	"math"

	"github.com/aliyun/aliyun-tablestore-go-sdk/timeline"
	"github.com/aliyun/aliyun-tablestore-go-sdk/timeline/promise"
	"github.com/aliyun/aliyun-tablestore-go-sdk/timeline/writer"
)

var IMJoiner = ":"

type IM struct {
	historyStore timeline.MessageStore
	syncStore    timeline.MessageStore

	adapter timeline.MessageAdapter
}

func NewIm(storeOption, syncOption timeline.StoreOption, adapter timeline.MessageAdapter) (*IM, error) {
	history, err := timeline.NewDefaultStore(storeOption)
	if err != nil {
		return nil, err
	}
	// if table is not exist, sync will create table
	// if table is already exist and StoreOption.TTL is not zero, sync will check and update table TTL if needed
	err = history.Sync()
	if err != nil {
		return nil, err
	}
	sync, err := timeline.NewDefaultStore(syncOption)
	if err != nil {
		return nil, err
	}
	err = sync.Sync()
	if err != nil {
		return nil, err
	}
	im := &IM{
		historyStore: history,
		syncStore:    sync,
		adapter:      adapter,
	}
	return im, nil
}

func (im *IM) GetSyncMessage(member string, lastRead int64) ([]*timeline.Entry, error) {
	receiver, err := timeline.NewTmLine(member, im.adapter, im.syncStore)
	if err != nil {
		return nil, err
	}
	iterator := receiver.Scan(&timeline.ScanParameter{
		From:        lastRead,
		To:          math.MaxInt64,
		IsForward:   true,
		MaxCount:    100,
		BufChanSize: 10,
	})
	entries := make([]*timeline.Entry, 0)
	//avoid scanner goroutine leak
	defer iterator.Close()
	for {
		entry, err := iterator.Next()
		if err != nil {
			if err == timeline.ErrorDone {
				break
			} else {
				return entries, err
			}
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func (im *IM) GetHistoryMessage(storeName string, numOfHistory int, lastRead int64) ([]*timeline.Entry, int64, error) {
	if lastRead <= 0 {
		lastRead = math.MaxInt64
	}
	var latestRead int64 = math.MaxInt64
	receiver, err := timeline.NewTmLine(storeName, im.adapter, im.historyStore)
	if err != nil {
		return nil, latestRead, err
	}

	iterator := receiver.Scan(&timeline.ScanParameter{
		From:        lastRead,
		To:          0,
		MaxCount:    numOfHistory + 1,
		BufChanSize: 15,
	})

	entries := make([]*timeline.Entry, 0)
	//avoid scanner goroutine leak
	defer iterator.Close()

	for index := 0; index <= numOfHistory; index++ {
		entry, err := iterator.Next()
		if index == numOfHistory {
			if err == timeline.ErrorDone {
				latestRead = 1
			} else {
				latestRead = entry.Sequence
			}
		}
		if err == timeline.ErrorDone {
			latestRead = 1
			break
		}
		entries = append(entries, entry)
	}
	return entries, latestRead, nil
}

func (im *IM) Send(from, to string, message timeline.Message) (seq1 int64, seq2 int64, err error) {
	sender, err := timeline.NewTmLine(SingChatStoreName(from, to), im.adapter, im.historyStore)
	if err != nil {
		return
	}
	seq1, err = sender.Store(message)
	if err != nil {
		return
	}

	receiver, err := timeline.NewTmLine(to, im.adapter, im.syncStore)
	if err != nil {
		return
	}
	seq2, err = receiver.Store(message)
	if err != nil {
		return
	}
	return
}

func (im *IM) SendGroup(groupName string, groupMembers []string, message timeline.Message) ([]string, error) {
	sender, err := timeline.NewTmLine(groupName, im.adapter, im.historyStore)
	if err != nil {
		return nil, err
	}
	seq, err := sender.Store(message)
	if err != nil {
		return nil, err
	}
	fmt.Println("message auto increment sequence", seq)

	futures := make([]*promise.Future, len(groupMembers))
	for i, m := range groupMembers {
		receiver, err := timeline.NewTmLine(m, im.adapter, im.syncStore)
		if err != nil {
			return nil, err
		}
		f, err := receiver.BatchStore(message)
		if err != nil {
			return nil, err
		}
		futures[i] = f
	}

	fanFuture := promise.FanIn(futures...)
	fanResult, err := fanFuture.FanInGet()
	if err != nil {
		return nil, err
	}
	failedId := make([]string, 0)
	for _, result := range fanResult {
		if result.Err != nil {
			failedId = append(failedId, result.Result.(*writer.BatchAddResult).Id)
		}
	}
	return failedId, nil
}

func (im *IM) Close() {
	im.syncStore.Close()
	im.historyStore.Close()
}

func SingChatStoreName(userA, userB string) string {
	if userA > userB {
		return userB + IMJoiner + userA
	}
	return userA + IMJoiner + userB
}
