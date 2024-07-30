package userservicelogic

import (
	"context"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUsersLogic {
	return &GetAllUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllUsersLogic) GetAllUsers(in *pb.GetAllUsersReq) (*pb.GetAllUsersResp, error) {
	// todo: add your logic here and delete this line
	selectBuilder := l.svcCtx.UserModel.SelectBuilder()
	users, err := l.svcCtx.UserModel.FindAll(l.ctx, selectBuilder, "")
	if err != nil {
		return nil, errors.Wrapf(err, "get users error:%s", err.Error())
	}
	resp := new(pb.GetAllUsersResp)
	for _, user := range users {
		var u pb.User
		copier.Copy(&u, user)
		resp.Users = append(resp.Users, &u)
	}
	return resp, nil
}
