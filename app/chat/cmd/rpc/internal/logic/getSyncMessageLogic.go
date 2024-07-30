package logic

import (
	"context"
	"strconv"

	"zero-chat/app/chat/cmd/rpc/internal/svc"
	"zero-chat/app/chat/cmd/rpc/pb"
	"zero-chat/app/chat/model"

	"github.com/aliyun/aliyun-tablestore-go-sdk/timeline"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSyncMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSyncMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSyncMessageLogic {
	return &GetSyncMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSyncMessageLogic) GetSyncMessage(in *pb.GetSyncMessageReq) (resp *pb.GetSyncMessageResp, err error) {
	// todo: add your logic here and delete this line
	entries, err := l.svcCtx.IM.GetSyncMessage("user_"+strconv.Itoa(int(in.UserId)), in.LastRead)
	if err != nil {
		return nil, errors.Wrapf(err, "get sync msg failed:%s", err.Error())
	}
	resp = new(pb.GetSyncMessageResp)
	for _, entry := range entries {
		smsg := entry.Message.(*timeline.StreamMessage)
		msg := pb.MessageWithSeq{
			Sequence:    entry.Sequence,
			Content:     smsg.Content.(string),
			From:        smsg.Attr[model.Sender].(int64),
			SendTime:    smsg.Timestamp,
			ContentType: int32(smsg.Attr[model.MsgType].(int64)),
			File:        smsg.Attr[model.File].(string),
			ChatType:    int32(smsg.Attr[model.Type].(int64)),
		}
		resp.Msg = append(resp.Msg, &msg)
	}
	return
}
