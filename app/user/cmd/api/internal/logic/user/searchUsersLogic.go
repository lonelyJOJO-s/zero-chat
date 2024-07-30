package user

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

type SearchUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUsersLogic {
	return &SearchUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchUsersLogic) SearchUsers(req *types.SearchUsersReq) (resp *types.SearchUsersResp, err error) {
	id := ctxdata.GetUidFromCtx(l.ctx)
	userResp, err := l.svcCtx.UserServiceRpc.SearchUserFuzzy(l.ctx, &pb.SearchUserFuzzyReq{Keyword: req.Keyword})
	if err != nil {
		return nil, errors.Wrapf(err, "get users error with:%s", err.Error())
	}
	friendResp, err := l.svcCtx.FriendServiceRpc.GetFriends(l.ctx, &pb.GetFriendsReq{Id: id})
	if err != nil {
		return nil, errors.Wrapf(err, "get user friends error with:%s", err.Error())
	}
	friendIds := map[int64]bool{id: true}
	for _, friend := range friendResp.Users {
		friendIds[friend.Id] = true
	}
	resp = new(types.SearchUsersResp)
	for _, user := range userResp.Users {
		if _, ok := friendIds[user.Id]; !ok {
			var u types.User
			copier.Copy(&u, user)
			resp.Users = append(resp.Users, u)
		}
	}
	return
}
