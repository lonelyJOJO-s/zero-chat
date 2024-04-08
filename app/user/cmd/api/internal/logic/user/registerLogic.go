package user

import (
	"context"

	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	var user pb.UserWithPwd
	copier.Copy(&user, req)
	registerResp, err := l.svcCtx.UserServiceRpc.Register(l.ctx, &pb.RegisterReq{UserInfo: &user})
	if err != nil {
		return nil, err
	}
	resp = new(types.RegisterResp)
	copier.Copy(resp, registerResp)
	return
}
