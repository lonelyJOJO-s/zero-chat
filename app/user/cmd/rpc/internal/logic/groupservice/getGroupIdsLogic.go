package groupservicelogic

import (
	"context"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/xerr"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupIdsLogic {
	return &GetGroupIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupIdsLogic) GetGroupIds(in *pb.GetGroupIdsReq) (*pb.GetGroupIdsResp, error) {
	// todo: add your logic here and delete this line
	selectBuild := l.svcCtx.GroupModel.SelectBuilder().Where(squirrel.Eq{"`owner_id`": in.UserId})
	ids, err := l.svcCtx.GroupModel.FindAllByUserId(l.ctx, selectBuild, "")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find groups error:%s", err.Error())
	}
	return &pb.GetGroupIdsResp{GroupIds: ids}, nil
}
