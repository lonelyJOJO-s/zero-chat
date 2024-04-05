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

type GetMemberIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMemberIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMemberIdsLogic {
	return &GetMemberIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMemberIdsLogic) GetMemberIds(in *pb.GetMemberIdsReq) (*pb.GetMemberIdsResp, error) {
	rowBuilder := l.svcCtx.UserGroup.SelectBuilder().Where(squirrel.Eq{"`group_id`": in.GroupId})
	ids, err := l.svcCtx.UserGroup.GetUserIdsByGroupId(l.ctx, rowBuilder, "")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find members id failed with error:%s", err.Error())
	}
	return &pb.GetMemberIdsResp{Ids: ids}, nil
}
