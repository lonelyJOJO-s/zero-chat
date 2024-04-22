package groupservicelogic

import (
	"context"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetJoinedGroupIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetJoinedGroupIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJoinedGroupIdsLogic {
	return &GetJoinedGroupIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetJoinedGroupIdsLogic) GetJoinedGroupIds(in *pb.GetJoinedGroupIdsReq) (*pb.GetJoinedGroupIdsResp, error) {
	ans, err := l.svcCtx.UserGroup.FindAllIdsByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get group ids error:%s", err.Error())
	}
	return &pb.GetJoinedGroupIdsResp{GroupIds: ans}, nil
}
