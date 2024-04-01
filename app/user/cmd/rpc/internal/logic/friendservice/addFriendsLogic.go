package friendservicelogic

import (
	"context"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendsLogic {
	return &AddFriendsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddFriendsLogic) AddFriends(in *pb.AddFriendsReq) (*pb.AddFriendsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddFriendsResp{}, nil
}
