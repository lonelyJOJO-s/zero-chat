package userservicelogic

import (
	"context"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/app/user/model"
	"zero-chat/common/xerr"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersInfoLogic {
	return &GetUsersInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// user basic
func (l *GetUsersInfoLogic) GetUsersInfo(in *pb.GetUsersInfoReq) (resp *pb.GetUsersInfoResp, err error) {
	selectBuilder := l.svcCtx.UserModel.SelectBuilder().Where(squirrel.Eq{"`id`": in.Ids})
	users, err := l.svcCtx.UserModel.FindAll(l.ctx, selectBuilder, "")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserInfo find user db err:%v", err)
	}
	if len(users) == 0 {
		return nil, errors.Wrapf(ErrUserNoExistsError, "ids:%v", in.Ids)
	}

	resp = new(pb.GetUsersInfoResp)
	for _, user := range users {
		var respUser pb.User
		copier.Copy(&respUser, user)
		resp.Users = append(resp.Users, &respUser)
	}

	return
}
