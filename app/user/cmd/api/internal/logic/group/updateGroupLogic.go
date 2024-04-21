package group

import (
	"context"

	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupLogic {
	return &UpdateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGroupLogic) UpdateGroup(req *types.GroupUpdateReq) (resp *types.Null, err error) {
	_, err = l.svcCtx.GroupServiceRpc.UpdateGroupInfo(l.ctx, &pb.UpdateGroupReq{
		Group: &pb.Group{
			Id:      req.Group.Id,
			Name:    req.Group.Name,
			Desc:    req.Group.Desc,
			Avatar:  req.Group.Avatar,
			OwnerId: req.Group.CreatorId,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "update group error with:%s", err.Error())
	}
	return
}
