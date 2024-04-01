package user

import (
	"context"

	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
