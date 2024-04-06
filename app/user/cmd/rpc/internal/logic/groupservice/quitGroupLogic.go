package groupservicelogic

import (
	"context"
	"slices"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
	group, err := l.svcCtx.GroupModel.FindOne(l.ctx, in.GroupId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find group error:%s", err.Error())
	}

	if in.UserId != group.OwnerId {
		err = l.svcCtx.UserGroup.DeleteSoft(l.ctx, userGroup.Id, nil)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "del relationship error:%s", err.Error())
		}
		return &pb.QuitGroupResp{}, nil
	}

	getMembersLogic := NewGetMemberIdsLogic(l.ctx, l.svcCtx)
	candidates, err := getMembersLogic.GetMemberIds(&pb.GetMemberIdsReq{GroupId: in.GroupId})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get members error:%s", err.Error())
	}
	logx.Infof("candidates:%v", candidates.Ids)
	// 没有指认继承者，则
	if in.HeirId == 0 {
		if len(candidates.Ids) > 1 {
			return nil, xerr.NewErrCode(xerr.MUST_CHOOSE_HEIR)
		} else {
			// del group (没有执行)
			err := l.svcCtx.GroupModel.DeleteSoft(l.ctx, in.GroupId)
			if err != nil {
				return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "del group error:%s", err.Error())
			}
			// del relatinship
			err = l.svcCtx.UserGroup.DeleteSoft(l.ctx, userGroup.Id, nil)
			if err != nil {
				return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "del relationship error:%s", err.Error())
			}
			return &pb.QuitGroupResp{}, nil
		}
	}
	if !slices.Contains(candidates.Ids, in.HeirId) {
		return nil, xerr.NewErrCode(xerr.UserNotInGroup)
	}
	// use transaction
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		err = l.svcCtx.UserGroup.DeleteSoft(l.ctx, userGroup.Id, session)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "del relationship error:%s", err.Error())
		}
		group.OwnerId = in.HeirId
		err = l.svcCtx.GroupModel.Update(l.ctx, group, session)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "update group info error:%s", err.Error())
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &pb.QuitGroupResp{}, nil

}
