package groupservicelogic

import (
	"context"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupInfoLogic {
	return &UpdateGroupInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupInfoLogic) UpdateGroupInfo(in *pb.UpdateGroupReq) (*pb.UpdateGroupResp, error) {
	update, err := l.svcCtx.GroupModel.FindOne(l.ctx, in.Group.Id)
	if err != nil {
		return nil, err
	}
	if update.OwnerId != in.Group.OwnerId {
		return nil, xerr.NewErrCode(xerr.NO_ACCESS_TO_RESOURCE)
	}
	copier.CopyWithOption(&update, in.Group, copier.Option{IgnoreEmpty: true})
	err = l.svcCtx.GroupModel.Update(l.ctx, update, nil)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "update group info error:%s", err.Error())
	}
	return &pb.UpdateGroupResp{}, nil
}
