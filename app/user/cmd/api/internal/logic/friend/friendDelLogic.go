package friend

import (
	"context"

	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendDelLogic {
	return &FriendDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendDelLogic) FriendDel(req *types.FriendIdReq) (resp *types.Null, err error) {
	// todo: add your logic here and delete this line

	return
}
