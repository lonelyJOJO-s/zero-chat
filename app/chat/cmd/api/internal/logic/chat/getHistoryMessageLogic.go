package chat

import (
	"context"

	"zero-chat/app/chat/cmd/api/internal/svc"
	"zero-chat/app/chat/cmd/api/internal/types"
	"zero-chat/app/chat/cmd/rpc/pb"
	userPb "zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/constant"
	"zero-chat/common/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetHistoryMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHistoryMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHistoryMessageLogic {
	return &GetHistoryMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHistoryMessageLogic) GetHistoryMessage(req *types.GetHistoryMessageReq) (resp *types.GetHistoryMessageResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	var messages *pb.GetHistoryMessageResp
	switch req.ChatType {
	case constant.SINGLE:
		messages, err = l.svcCtx.ChatServiceRpc.GetHistoryMessage(l.ctx, &pb.GetHistoryMessageReq{
			UserA: int64(req.Id), UserB: userId, Nums: req.Cnt})
	case constant.GROUP:
		messages, err = l.svcCtx.ChatServiceRpc.GetHistoryMessage(l.ctx, &pb.GetHistoryMessageReq{
			GroupId: int64(req.Id), Nums: req.Cnt})
	}
	if err != nil {
		return nil, errors.Wrapf(err, "api--- get history msg error: %s", err.Error())
	}
	resp = new(types.GetHistoryMessageResp)
	for _, msg := range messages.Msg {
		var m types.Message
		user, _ := l.svcCtx.UserServiceRpc.GetUsersInfo(l.ctx, &userPb.GetUsersInfoReq{Ids: []int64{msg.From}})
		copier.Copy(&m, msg)
		m.Avatar = user.Users[0].Avatar
		m.FromUsername = user.Users[0].Username
		resp.Msgs = append(resp.Msgs, m)
	}
	return
}
