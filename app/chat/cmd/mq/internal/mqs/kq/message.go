package kq

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"zero-chat/app/chat/cmd/mq/internal/svc"
	"zero-chat/app/chat/cmd/rpc/pb"
	userPb "zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/constant"
	"zero-chat/common/protocol"
	"zero-chat/common/tool"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"google.golang.org/protobuf/proto"
)

/*
*
Listening to the payment flow status change notification message queue
*/
type MessageTransferMq struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

const (
	cacheMQUsersIdPrefix = "cache:mq:user_id:"
)

func NewMessageTransferMq(ctx context.Context, svcCtx *svc.ServiceContext) *MessageTransferMq {
	return &MessageTransferMq{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageTransferMq) Consume(_, val string) error {

	var message protocol.Message
	if err := proto.Unmarshal([]byte(val), &message); err != nil {
		logx.WithContext(l.ctx).Errorf("MessageTransferMq->Consume Unmarshal err : %v , val : %s", err, val)
		return err
	}
	if err := l.execService(val, &message); err != nil {
		logx.Errorf("MessageTransferMq->execService  err : %v , val : %s", err, val)
		return err
	}

	return nil
}

func (l *MessageTransferMq) execService(val string, message *protocol.Message) (err error) {

	var receivers []int64

	// 1. store into store-table (session —— read expand)
	storeMsg := pb.StoreTableItem{
		TimeLineId:  uuid.NewString(),
		SequenceId:  time.Now().UnixNano(),
		MsgType:     message.ChatType,
		ContentType: message.ContentType,
		SendTime:    message.SendTime,
		Sender:      message.From,
		Content:     message.Content,
		File:        message.File,
	}

	switch message.ChatType {
	case constant.SINGLE:
		resp, err := l.svcCtx.FriendServiceRpc.GetUuid(l.ctx, &userPb.GetUuidReq{UserId: message.From, FriendId: message.To})
		if err != nil {
			return errors.Wrapf(err, "get relation uuid error:%s", err.Error())
		}
		storeMsg.Conversation = string(resp.Uuid)
		receivers = []int64{message.To}
	case constant.GROUP:
		storeMsg.Conversation = strconv.Itoa(int(message.To))
		resp, err := l.svcCtx.GroupServiceRpc.GetMemberIds(l.ctx, &userPb.GetMemberIdsReq{GroupId: message.To})
		if err != nil {
			return errors.Wrapf(err, "get group ids error:%s", err.Error())
		}
		receivers = resp.Ids

	}
	_, err = l.svcCtx.ChatServiceRpc.StoreAddItem(l.ctx, &pb.StoreAddItemReq{Msg: &storeMsg})
	if err != nil {
		return errors.Wrapf(err, "insert store-table error:%s", err.Error())
	}
	// 2. store into sync-table, (mail —— write expand)
	syncMsg := pb.SyncTableItem{
		MsgType:     message.ChatType,
		ContentType: message.ContentType,
		SendTime:    message.SendTime,
		Sender:      message.From,
		Content:     message.Content,
	}
	for _, recevier := range receivers {
		syncMsg.UserId = recevier
		syncMsg.TimeLineId = uuid.NewString()
		syncMsg.SequenceId = time.Now().UnixNano()
		_, err = l.svcCtx.ChatServiceRpc.SyncAddItem(l.ctx, &pb.SyncAddItemReq{Msg: &syncMsg})
		if err != nil {
			return err
		}
		// 3. store into redis, use the distribute lock
		lock := redis.NewRedisLock(l.svcCtx.Redis, fmt.Sprintf("%s%d:%s", constant.DISTRIBUTE_PREFIX, syncMsg.Sender, storeMsg.Conversation))
		acquire, err := lock.Acquire()
		var retries int = 0
	FOR:
		for retries < constant.MAX_RETRY {
			switch {
			case err != nil:
				return errors.Wrapf(err, "get reids lock err:%s", err.Error())
			case acquire:
				// TODO: 添加事务支持
				defer lock.Release() // 释放锁
				key := fmt.Sprintf("%s%d:conversation:%s", cacheMQUsersIdPrefix, syncMsg.UserId, storeMsg.Conversation)

				l.svcCtx.Redis.HsetCtx(l.ctx, key, "lastest_content", message.Content)

				_, err := l.svcCtx.Redis.HgetCtx(l.ctx, key, "sequence_id")
				switch {
				case err == redis.Nil:
					l.svcCtx.Redis.HsetCtx(l.ctx, key, "sequence_id", fmt.Sprint(syncMsg.SequenceId))
				case err != nil:
					return errors.Wrapf(err, "get cnt of user:%d err:%s", message.From, err.Error())
				}

				cnt, err := l.svcCtx.Redis.HgetCtx(l.ctx, key, "cnt")
				switch {
				case err == redis.Nil:
					l.svcCtx.Redis.HsetCtx(l.ctx, key, "cnt", "1")
				case err != nil:
					return errors.Wrapf(err, "get cnt of user:%d err:%s", message.From, err.Error())
				default:
					cnt, _ = tool.StrAutoIncrease(cnt)
					l.svcCtx.Redis.HsetCtx(l.ctx, key, "cnt", cnt)

				}
				break FOR
			case !acquire:
				time.Sleep(constant.RETRY_INTERVAL * time.Millisecond)
				retries++
			}

		}

	}
	err = l.svcCtx.KqPusherClient.Push(val)
	if err != nil {
		return errors.Wrapf(err, "send to msg back to kafka error:%s", err.Error())
	}
	return err
}
