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

type SearchHistoryMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchHistoryMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchHistoryMessageLogic {
	return &SearchHistoryMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchHistoryMessageLogic) SearchHistoryMessage(req *types.SearchHistoryMessageReq) (resp *types.SearchHistoryMessageResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	var messages *pb.SearchHistoryMessageResp
	switch req.ChatType {
	case constant.SINGLE:
		messages, err = l.svcCtx.ChatServiceRpc.SearchHistoryMEssage(l.ctx, &pb.SearchHistoryMessageReq{
			UserA: int64(req.Id), UserB: userId, Offset: req.Offset, Limit: req.Limit, KeyWord: req.Keyword})
	case constant.GROUP:
		messages, err = l.svcCtx.ChatServiceRpc.SearchHistoryMEssage(l.ctx, &pb.SearchHistoryMessageReq{
			GroupId: int64(req.Id), Offset: req.Offset, Limit: req.Limit, KeyWord: req.Keyword})
	}
	if err != nil {
		return nil, errors.Wrapf(err, "api--- get history msg error: %s", err.Error())
	}
	resp = new(types.SearchHistoryMessageResp)
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
