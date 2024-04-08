package user

import (
	"context"
	"fmt"

	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWithUsernameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginWithUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithUsernameLogic {
	return &LoginWithUsernameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginWithUsernameLogic) LoginWithUsername(req *types.UsernameLoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	pbResp, err := l.svcCtx.UserServiceRpc.Login(l.ctx, &pb.LoginReq{
		Type:     1,
		Account:  req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginResp)
	copier.Copy(resp, pbResp)
	fmt.Println(resp)
	return
}
