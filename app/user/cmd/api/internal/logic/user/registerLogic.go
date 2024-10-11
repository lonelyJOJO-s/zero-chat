package user

import (
	"context"
	"fmt"
	"os"

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

var (
	url          = os.Getenv("OSS_URI")
	AvatarMale   = fmt.Sprintf("%s/avatar/男.jpg", url)
	AvatarFamale = fmt.Sprintf("%s/avatar/女.jpg", url)
)

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
	if user.Sex == 0 {
		user.Avatar = AvatarMale
	} else {
		user.Avatar = AvatarFamale
	}
	registerResp, err := l.svcCtx.UserServiceRpc.Register(l.ctx, &pb.RegisterReq{UserInfo: &user})
	if err != nil {
		return nil, err
	}
	resp = new(types.RegisterResp)
	copier.Copy(resp, registerResp)
	return
}
