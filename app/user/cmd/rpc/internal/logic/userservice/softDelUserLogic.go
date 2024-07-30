package userservicelogic

import (
	"context"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SoftDelUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSoftDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SoftDelUserLogic {
	return &SoftDelUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SoftDelUserLogic) SoftDelUser(in *pb.DelUserInfoReq) (*pb.DelUserInfoResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.UserModel.Delete(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.DelUserInfoResp{Code: 0}, nil
}
