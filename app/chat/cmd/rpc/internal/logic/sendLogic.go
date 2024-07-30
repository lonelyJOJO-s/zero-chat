package logic

import (
	"context"
	"strconv"
	"time"

	"zero-chat/app/chat/cmd/rpc/internal/svc"
	"zero-chat/app/chat/cmd/rpc/pb"
	"zero-chat/app/chat/model"
	userPb "zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/constant"

	"github.com/aliyun/aliyun-tablestore-go-sdk/timeline"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendLogic) Send(in *pb.SendReq) (resp *pb.SendResp, err error) {

	// igonre sync table's id field
	var seq1, seq2 int64
	var id string
	attr := map[string]interface{}{
		model.MsgType: in.Msg.ContentType,
		model.Type:    in.Msg.ChatType,
		model.Sender:  in.Msg.From,
		model.File:    in.Msg.File,
	}
	t := time.Now().UnixNano()
	msg := &timeline.StreamMessage{
		Content:   in.Msg.Content,
		Timestamp: t,
		Attr:      attr,
	}
	switch in.Msg.ChatType {
	case constant.SINGLE:
		idInt, err := l.svcCtx.FriendServiceRpc.GetUuid(l.ctx, &userPb.GetUuidReq{UserId: in.Msg.From, FriendId: in.Msg.To})
		if err != nil {
			return nil, errors.Wrapf(err, "get users relation failed:%s", err.Error())
		}
		id = idInt.Uuid
		msg.Id = id
		seq1, seq2, err = l.svcCtx.IM.Send("user_"+strconv.Itoa(int(in.Msg.From)), "user_"+strconv.Itoa(int(in.Msg.To)), msg)
		if err != nil {
			return nil, errors.Wrapf(err, "send im msg failed:%s", err.Error())
		}
	case constant.GROUP:
		ids, err := l.svcCtx.GroupServiceRpc.GetMemberIds(l.ctx, &userPb.GetMemberIdsReq{GroupId: in.Msg.To})
		if err != nil {
			return nil, errors.Wrapf(err, "get group members failed:%s", err.Error())
		}
		members := []string{}
		for _, id := range ids.Ids {
			members = append(members, "user_"+strconv.Itoa(int(id)))
		}
		msg.Id = strconv.Itoa(int(in.Msg.To))
		failedIds, err := l.svcCtx.IM.SendGroup("group_"+strconv.Itoa(int(in.Msg.To)), members, msg)
		if err != nil {
			return nil, errors.Wrapf(err, "send group failed:%s, failedIds:%v", err.Error(), failedIds)
		}
	}
	return &pb.SendResp{StoreSequence: seq1, SyncSequence: seq2}, nil
}
