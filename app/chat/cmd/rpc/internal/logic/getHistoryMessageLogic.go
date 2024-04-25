package logic

import (
	"context"
	"strconv"

	"zero-chat/app/chat/cmd/rpc/internal/im"
	"zero-chat/app/chat/cmd/rpc/internal/svc"
	"zero-chat/app/chat/cmd/rpc/pb"
	"zero-chat/app/chat/model"

	"github.com/aliyun/aliyun-tablestore-go-sdk/timeline"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetHistoryMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetHistoryMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHistoryMessageLogic {
	return &GetHistoryMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetHistoryMessageLogic) GetHistoryMessage(in *pb.GetHistoryMessageReq) (resp *pb.GetHistoryMessageResp, err error) {
	var entries []*timeline.Entry
	if in.GroupId > 0 {
		entries, err = l.svcCtx.IM.GetHistoryMessage("group_"+strconv.Itoa(int(in.GroupId)), int(in.Nums))
	} else {
		timelineId := im.SingChatStoreName("user_"+strconv.Itoa(int(in.UserA)), "user_"+strconv.Itoa(int(in.UserB)))
		entries, err = l.svcCtx.IM.GetHistoryMessage(timelineId, int(in.Nums))
	}
	if err != nil {
		return nil, errors.Wrapf(err, "get histpry msgs failed:%s", err.Error())
	}
	resp = new(pb.GetHistoryMessageResp)
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
