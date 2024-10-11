package group

import (
	"context"

	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/ctxdata"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllGroupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllGroupsLogic {
	return &GetAllGroupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllGroupsLogic) GetAllGroups(req *types.Null) (resp *types.GetAllGroupsResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	groupResp, err := l.svcCtx.GroupServiceRpc.GetJoinedGroups(l.ctx, &pb.GetJoinedGroupsReq{UserId: userId})
	if err != nil {
		return nil, errors.Wrapf(err, "get groups error with:%s", err.Error())
	}
	resp = new(types.GetAllGroupsResp)
	for _, group := range groupResp.Groups {
		g := types.GroupWithId{
			Id: group.Id,
			Group: types.Group{
				Name:            group.Name,
				Desc:            group.Desc,
				Avatar:          group.Avatar,
				CreatorId:       group.OwnerId,
				LastMessageTime: group.LastMessageTime,
			},
		}
		resp.Groups = append(resp.Groups, g)
	}
	return
}
