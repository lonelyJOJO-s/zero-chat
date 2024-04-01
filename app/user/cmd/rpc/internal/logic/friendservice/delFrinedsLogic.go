package friendservicelogic

import (
	"context"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelFrinedsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelFrinedsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFrinedsLogic {
	return &DelFrinedsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelFrinedsLogic) DelFrineds(in *pb.DelFriendsReq) (*pb.DelFriendsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelFriendsResp{}, nil
}
