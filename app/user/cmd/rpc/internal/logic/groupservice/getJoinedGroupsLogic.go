package groupservicelogic

import (
	"context"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/xerr"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetJoinedGroupsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetJoinedGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJoinedGroupsLogic {
	return &GetJoinedGroupsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetJoinedGroupsLogic) GetJoinedGroups(in *pb.GetJoinedGroupsReq) (resp *pb.GetJoinedGroupsResp, err error) {

	getJoinedGroupIdsLogic := NewGetJoinedGroupIdsLogic(l.ctx, l.svcCtx)
	idsResp, err := getJoinedGroupIdsLogic.GetJoinedGroupIds(&pb.GetJoinedGroupIdsReq{UserId: in.UserId})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get group ids error:%s", err.Error())
	}
	selectBuilder := l.svcCtx.GroupModel.SelectBuilder().Where(squirrel.Eq{"`id`": idsResp.GroupIds})
	groups, err := l.svcCtx.GroupModel.FindAll(l.ctx, selectBuilder, "")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get groups error:%s", err.Error())
	}
	resp = new(pb.GetJoinedGroupsResp)
	for _, group := range groups {
		var g pb.Group
		copier.Copy(&g, group)
		resp.Groups = append(resp.Groups, &g)
	}
	return resp, nil
}
