package chat

import (
	"context"

	"zero-chat/app/chat/cmd/api/internal/svc"
	"zero-chat/app/chat/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SessionContentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSessionContentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SessionContentListLogic {
	return &SessionContentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SessionContentListLogic) SessionContentList(req *types.SessionContentListReq) (resp *types.SessionContentListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
