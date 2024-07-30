package friendservicelogic

import (
	"context"

	"github.com/pkg/errors"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelFrinedsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelFrinedsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFrinedsLogic {
	return &DelFrinedsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelFrinedsLogic) DelFrineds(in *pb.DelFriendsReq) (*pb.DelFriendsResp, error) {
	uf, err := l.svcCtx.UserFriend.FindOneByUserIdFriendId(l.ctx, in.Id, in.FriendId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find userfrined error with:%s", err.Error())
	}
	err = l.svcCtx.UserFriend.DeleteSoft(l.ctx, *uf)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "del frined error with:%s", err.Error())
	}
	return &pb.DelFriendsResp{}, nil
}
