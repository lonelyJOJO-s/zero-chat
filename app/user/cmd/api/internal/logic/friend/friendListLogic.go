package friend

import (
	"context"

	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.Null) (resp *types.FriendsResp, err error) {
	id := ctxdata.GetUidFromCtx(l.ctx)
	friendsResp, err := l.svcCtx.FriendServiceRpc.GetFriends(l.ctx, &pb.GetFriendsReq{Id: id})
	if err != nil {
		return nil, errors.Wrapf(err, "del friend rpc error with:%s", err.Error())
	}
	resp = new(types.FriendsResp)
	for _, user := range friendsResp.Users {
		u := types.User{}
		copier.Copy(&u, user)
		resp.Users = append(resp.Users, u)
	}
	return
}
