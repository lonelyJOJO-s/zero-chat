package groupservicelogic

import (
	"context"
	"slices"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type QuitGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuitGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuitGroupLogic {
	return &QuitGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QuitGroupLogic) QuitGroup(in *pb.QuitGroupReq) (*pb.QuitGroupResp, error) {
	// 需要保证user_id 和 heir_id不同，如果只剩一个用户，则输入heir_id = 0
	userGroup, err := l.svcCtx.UserGroup.FindOneByUserIdGroupId(l.ctx, in.UserId, in.GroupId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find relationship error:%s", err.Error())
	}
	err = l.svcCtx.UserGroup.DeleteSoft(l.ctx, userGroup.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "del relationship error:%s", err.Error())
	}
	group, err := l.svcCtx.GroupModel.FindOne(l.ctx, in.GroupId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find group error:%s", err.Error())
	}
	if in.UserId != group.OwnerId {
		return &pb.QuitGroupResp{}, nil
	}
	getMembersLogic := NewGetMemberIdsLogic(l.ctx, l.svcCtx)
	candidates, err := getMembersLogic.GetMemberIds(&pb.GetMemberIdsReq{GroupId: in.GroupId})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get members error:%s", err.Error())
	}
	if !slices.Contains(candidates.Ids, in.HeirId) {
		return nil, xerr.NewErrCode(xerr.UserNotInGroup)
	}
	updateGroupInfoLogic := NewUpdateGroupInfoLogic(l.ctx, l.svcCtx)
	var pbGroup pb.Group
	group.OwnerId = in.HeirId
	copier.Copy(&pbGroup, group)
	_, err = updateGroupInfoLogic.UpdateGroupInfo(&pb.UpdateGroupReq{Group: &pbGroup})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "update group info error:%s", err.Error())
	}
	return &pb.QuitGroupResp{}, nil
}
