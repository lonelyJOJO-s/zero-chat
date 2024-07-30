package user

import (
	"context"

	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LogWithEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogWithEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogWithEmailLogic {
	return &LogWithEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogWithEmailLogic) LogWithEmail(req *types.EmailLoginReq) (resp *types.LoginResp, err error) {
	pbResp, err := l.svcCtx.UserServiceRpc.Login(l.ctx, &pb.LoginReq{
		Type:     2,
		Account:  req.Email,
		Password: req.Captcha,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginResp)
	copier.Copy(resp, pbResp)
	return
}
