package group

import (
	"context"

	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupInfoLogic {
	return &GroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupInfoLogic) GroupInfo(req *types.GroupId) (resp *types.GroupResp, err error) {
	// todo: add your logic here and delete this line
	groupInfo, err := l.svcCtx.GroupServiceRpc.GetGroupInfo(l.ctx, &pb.GetGroupInfoReq{Id: req.Id})
	if err != nil {
		return nil, errors.Wrapf(err, "get group info error with:%s", err.Error())
	}
	resp = &types.GroupResp{
		Group: types.GroupWithId{
			Id: groupInfo.Group.Id,
			Group: types.Group{
				Name:      groupInfo.Group.Name,
				Desc:      groupInfo.Group.Desc,
				Avatar:    groupInfo.Group.Avatar,
				CreatorId: groupInfo.Group.OwnerId,
			},
		},
	}
	return
}
