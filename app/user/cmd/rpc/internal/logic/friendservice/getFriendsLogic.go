package friendservicelogic

import (
	"context"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/tool"
	"zero-chat/common/xerr"

	sq "github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendsLogic {
	return &GetFriendsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// friend basic
func (l *GetFriendsLogic) GetFriends(in *pb.GetFriendsReq) (*pb.GetFriendsResp, error) {
	friendIdsWithTime, err := l.svcCtx.UserFriend.GetFriendIdsAndLastTime(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "user_id:%d get friend ids error:%s", in.Id, err.Error())
	}
	friendIds := make([]int64, len(friendIdsWithTime))
	for i, item := range friendIdsWithTime {
		friendIds[i] = item.Id
	}
	logx.Info(friendIdsWithTime)
	whereBuilder := l.svcCtx.UserModel.SelectBuilder().Where(
		sq.Eq{"`id`": friendIds})
	users, err := l.svcCtx.UserModel.FindAll(l.ctx, whereBuilder, "")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "user_id:%d get friend error:%s", in.Id, err.Error())
	}
	var resp pb.GetFriendsResp
	copier.Copy(&resp.Users, &users)
	for i, item := range friendIdsWithTime {
		resp.Users[i].LastMessageTime = tool.Time2TimeStamp(item.LastTime)
	}
	logx.Info(resp.Users)
	return &resp, nil
}
