package group

import (
	"context"

	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupLogic) CreateGroup(req *types.GroupCreateReq) (resp *types.GroupResp, err error) {
	// todo: add your logic here and delete this line

	pbResp, err := l.svcCtx.GroupServiceRpc.CreateGroup(l.ctx, &pb.CreateGroupReq{Keyword: &pb.Group{
		Name:    req.GroupInfo.Name,
		Desc:    req.GroupInfo.Desc,
		Avatar:  req.GroupInfo.Avatar,
		OwnerId: req.GroupInfo.CreatorId,
	}})
	if err != nil {
		return nil, errors.Wrapf(err, "create group error with:%s", err.Error())
	}
	resp = &types.GroupResp{
		Group: types.GroupWithId{
			Id: pbResp.Id,
			Group: types.Group{
				Name:      req.GroupInfo.Name,
				Desc:      req.GroupInfo.Desc,
				Avatar:    req.GroupInfo.Avatar,
				CreatorId: req.GroupInfo.CreatorId,
			},
		},
	}
	return
}
