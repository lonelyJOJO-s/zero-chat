package friendservicelogic

import (
	"context"
	"strings"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFriendFuzzyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchFriendFuzzyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFriendFuzzyLogic {
	return &SearchFriendFuzzyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchFriendFuzzyLogic) SearchFriendFuzzy(in *pb.SearchFriendFuzzyReq) (*pb.SearchFriendFuzzyResp, error) {
	// search friends by email or phone or username

	getFriendsLogic := NewGetFriendsLogic(l.ctx, l.svcCtx)
	friends, err := getFriendsLogic.GetFriends(&pb.GetFriendsReq{Id: in.Id})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get friends error with:%s", err.Error())
	}
	var resp pb.SearchFriendFuzzyResp
	for _, friend := range friends.Users {
		if strings.Contains(strings.Split(friend.Email, "@")[0], in.Keyword) ||
			strings.Contains(friend.Phone, in.Keyword) ||
			strings.Contains(friend.Username, in.Keyword) {
			resp.Users = append(resp.Users, friend)
		}
	}
	return &resp, nil
}
