package user

import (
	"context"

	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllLogic {
	return &GetAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllLogic) GetAll(req *types.Null) (resp *types.GetAllResp, err error) {
	// todo: add your logic here and delete this line
	usersResp, err := l.svcCtx.UserServiceRpc.GetAllUsers(l.ctx, &pb.GetAllUsersReq{})
	if err != nil {
		return nil, errors.Wrapf(err, "get users from rpc to api error with:%s", err.Error())
	}
	resp = new(types.GetAllResp)
	for _, user := range usersResp.Users {
		var u types.User
		copier.Copy(&u, user)
		resp.Users = append(resp.Users, u)
	}
	return
}
