package kq

import (
	"context"
	"time"
	"zero-chat/app/chat/cmd/mq/internal/svc"
	"zero-chat/app/chat/cmd/rpc/pb"
	userPb "zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/protocol"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
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
	if err := l.execService(&message); err != nil {
		logx.WithContext(l.ctx).Errorf("MessageTransferMq->execService  err : %v , val : %s , message:%+v", err, val, message)
		return err
	}

	return nil
}

func (l *MessageTransferMq) execService(message *protocol.Message) (err error) {

	resp, err := l.svcCtx.FriendServiceRpc.GetUuid(l.ctx, &userPb.GetUuidReq{UserId: message.From, FriendId: message.To})
	if err != nil {
		return err
	}
	// 1. store into store-table (session —— read expand)
	storeMsg := pb.StoreTableItem{
		TimeLineId:   uuid.NewString(),
		SequenceId:   time.Now().UnixNano(),
		Conversation: string(resp.Uuid),
		MsgType:      message.ChatType,
		ContentType:  message.ContentType,
		SendTime:     message.SendTime,
		Sender:       message.From,
		Content:      message.Content,
		File:         message.File,
	}
	_, err = l.svcCtx.ChatServiceRpc.StoreAddItem(l.ctx, &pb.StoreAddItemReq{Msg: &storeMsg})
	if err != nil {
		return err
	}
	// 2. store into sync-table, (mail —— write expand)
	syncMsg := pb.SyncTableItem{
		TimeLineId:  uuid.NewString(),
		SequenceId:  time.Now().UnixNano(),
		UserId:      message.To,
		MsgType:     message.ChatType,
		ContentType: message.ContentType,
		SendTime:    message.SendTime,
		Sender:      message.From,
		Content:     message.Content,
	}
	_, err = l.svcCtx.ChatServiceRpc.SyncAddItem(l.ctx, &pb.SyncAddItemReq{Msg: &syncMsg})
	if err != nil {
		return err
	}
	// 3. store into redis

	// 4. store into redis sub/pub
	return err
}
