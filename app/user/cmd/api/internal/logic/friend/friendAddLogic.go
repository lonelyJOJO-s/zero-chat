package friend

import (
	"context"

	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/ctxdata"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FriendAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendAddLogic {
	return &FriendAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendAddLogic) FriendAdd(req *types.FriendIdReq) (resp *types.Null, err error) {
	// todo: add your logic here and delete this line
	id := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.FriendServiceRpc.AddFriends(l.ctx, &pb.AddFriendsReq{FriendId: req.Id, Id: id})
	if err != nil {
		return nil, errors.Wrapf(err, "add friend rpc error with:%s", err.Error())
	}
	return
}
