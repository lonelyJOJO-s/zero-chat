package im

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore/search"
	"github.com/aliyun/aliyun-tablestore-go-sdk/timeline"
	"github.com/aliyun/aliyun-tablestore-go-sdk/timeline/promise"
	"github.com/aliyun/aliyun-tablestore-go-sdk/timeline/writer"
	"github.com/pkg/errors"
)

var IMJoiner = ":"

type IM struct {
	historyStore timeline.MessageStore
	syncStore    timeline.MessageStore
	tableStore   *tablestore.TableStoreClient
	adapter      timeline.MessageAdapter
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
	ts := tablestore.NewClient(storeOption.Endpoint, storeOption.Instance, storeOption.AkId, storeOption.AkSecret)
	im := &IM{
		historyStore: history,
		syncStore:    sync,
		adapter:      adapter,
		tableStore:   ts,
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

func (im *IM) QueryWithPrimaryKeyAndContent(tableName, indexName, keyVal, keyWord string, limit, offset int32) (datas []CustomData, err error) {
	// 构建 BoolQuery
	boolQuery := &search.BoolQuery{}

	// 主键查询
	primaryKeyQuery := &search.TermQuery{
		FieldName: "TimelineId",
		Term:      keyVal, // 你的主键值
	}
	boolQuery.MustQueries = append(boolQuery.MustQueries, primaryKeyQuery)

	// 模糊匹配 Content 字段，类似 SQL 中的 LIKE '%yyy%'
	contentWildcardQuery := &search.WildcardQuery{
		FieldName: "Content",
		Value:     "*" + keyWord + "*", // 模糊匹配值
	}
	boolQuery.MustQueries = append(boolQuery.MustQueries, contentWildcardQuery)

	// 创建搜索请求
	searchRequest := &tablestore.SearchRequest{}
	searchRequest.SetTableName(tableName)
	searchRequest.SetIndexName(indexName)

	searchQuery := search.NewSearchQuery()
	searchQuery.SetQuery(boolQuery)
	searchQuery.SetLimit(limit)
	searchQuery.SetOffset(offset)
	searchQuery.SetSort(&search.Sort{
		[]search.Sorter{
			&search.PrimaryKeySort{
				Order: search.SortOrder_DESC.Enum(),
			},
		},
	})

	searchRequest.SetSearchQuery(searchQuery)
	searchRequest.SetColumnsToGet(&tablestore.ColumnsToGet{
		ReturnAll: true,
	})

	// 执行查询
	searchResponse, err := im.tableStore.Search(searchRequest)
	if err != nil {
		err = errors.Wrapf(err, "Search tableStore error: %v", err.Error())
		return
	}
	datas = make([]CustomData, len(searchResponse.Rows))
	for i, row := range searchResponse.Rows {
		jsonData, err := json.Marshal(row.Columns)
		if err != nil {
			return nil, err
		}
		var data CustomData
		err = json.Unmarshal([]byte(jsonData), &data)
		if err != nil {
			return nil, err
		}
		datas[i] = data
	}
	return
}

type Column struct {
	ColumnName string      `json:"ColumnName"`
	Value      interface{} `json:"Value"`
}

type CustomData struct {
	File        string `column:"Attr_File"`
	From        int64  `column:"Attr_Sender"`
	SendTime    int64  `column:"Timestamp"`
	Content     string `column:"Content"`
	ContentType int64  `column:"Attr_MsgType"`
	ChatType    int64  `column:"Attr_Type"`
}

// 自定义的 UnmarshalJSON 方法
func (cd *CustomData) UnmarshalJSON(data []byte) error {
	// 解析为 Column 的 slice
	var columns []Column
	if err := json.Unmarshal(data, &columns); err != nil {
		return err
	}

	// 通过反射获取 CustomData 的字段信息
	v := reflect.ValueOf(cd).Elem()
	t := v.Type()

	// 遍历所有 Column 并匹配对应的 struct 标签
	for _, column := range columns {
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get("column")
			if tag == column.ColumnName {
				// 根据字段类型来设置值
				switch v.Field(i).Kind() {
				case reflect.String:
					if val, ok := column.Value.(string); ok {
						v.Field(i).SetString(val)
					}
				case reflect.Int32, reflect.Int64:
					// 处理数字类型（JSON 中数字默认是 float64）
					if val, ok := column.Value.(float64); ok {
						v.Field(i).SetInt(int64(val))
					}
				}
			}
		}
	}
	return nil
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
