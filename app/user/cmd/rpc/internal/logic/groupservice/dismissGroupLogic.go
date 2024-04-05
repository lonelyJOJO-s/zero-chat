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

type DismissGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDismissGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DismissGroupLogic {
	return &DismissGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DismissGroupLogic) DismissGroup(in *pb.DismissGroupReq) (*pb.DismissGroupResp, error) {
	group, err := l.svcCtx.GroupModel.FindOne(l.ctx, in.GroupId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find group error:%s", err.Error())
	}
	if group.OwnerId != in.UserId {
		return nil, xerr.NewErrCode(xerr.NO_ACCESS_TO_RESOURCE)
	}
	// soft del group
	err = l.svcCtx.GroupModel.DeleteSoft(l.ctx, in.GroupId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "delete group error:%s", err.Error())
	}
	// del all the relation
	builder := l.svcCtx.UserFriend.UpdateBuilder().Where(squirrel.Eq{"`group_id`": in.GroupId})
	l.svcCtx.UserGroup.DelAllRelationByGroupId(l.ctx, builder, "")
	return &pb.DismissGroupResp{}, nil
}
