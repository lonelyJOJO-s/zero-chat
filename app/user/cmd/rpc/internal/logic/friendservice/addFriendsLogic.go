package friendservicelogic

import (
	"context"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/app/user/model"
	"zero-chat/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendsLogic {
	return &AddFriendsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddFriendsLogic) AddFriends(in *pb.AddFriendsReq) (*pb.AddFriendsResp, error) {

	err := l.svcCtx.UserFriend.Insert(l.ctx, &model.UserFriend{UserId: in.Id, FriendId: in.FriendId})
	if err != nil {
		return nil, err
	}
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.FriendId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find user db error:%s with user_id:%d", err, in.FriendId)
	}
	var userCopy pb.User
	copier.Copy(&userCopy, user)
	return &pb.AddFriendsResp{Users: &userCopy}, nil
}
