package groupservicelogic

import (
	"context"
	"strings"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/xerr"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchGroupLogic {
	return &SearchGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchGroupLogic) SearchGroup(in *pb.SearchGroupReq) (*pb.SearchGroupResp, error) {
	// todo: search all groups belong to user_id by keyword
	// 1. 找到所有的groups， 然后对比keyword和group name
	getGroupsLogic := NewGetJoinedGroupIdsLogic(l.ctx, l.svcCtx)
	idsResp, err := getGroupsLogic.GetJoinedGroupIds(&pb.GetJoinedGroupIdsReq{UserId: in.UserId})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get group ids error:%s", err.Error())
	}
	rowBuilder := l.svcCtx.GroupModel.SelectBuilder().Where(squirrel.Eq{"`id`": idsResp.GroupIds})
	groups, err := l.svcCtx.GroupModel.FindAll(l.ctx, rowBuilder, "")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get groups error:%s", err.Error())
	}
	var pbGroupList []*pb.Group
	for _, group := range groups {
		if strings.Contains(group.Name.String, in.Keyword) {
			var pbGroup pb.Group
			copier.Copy(&pbGroup, group)
			pbGroupList = append(pbGroupList, &pbGroup)
		}
	}
	return &pb.SearchGroupResp{Group: pbGroupList}, nil
}
