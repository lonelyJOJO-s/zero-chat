package groupservicelogic

import (
	"context"
	"time"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/app/user/model"
	"zero-chat/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// group
func (l *CreateGroupLogic) CreateGroup(in *pb.CreateGroupReq) (*pb.CreateGroupResp, error) {
	// TODO: why use var g *model.Group can't work
	var g model.Groups
	copier.CopyWithOption(&g, in.Keyword, copier.Option{IgnoreEmpty: true})
	now := time.Now()
	g.CreatedAt, g.UpdatedAt = now, now
	result, err := l.svcCtx.GroupModel.Insert(l.ctx, &g)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "insert group error:%s", err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get group id error:%s", err.Error())
	}
	joinGroupLogic := NewJoinGroupLogic(l.ctx, l.svcCtx)
	_, err = joinGroupLogic.JoinGroup(&pb.JoinGroupReq{UserId: in.Keyword.OwnerId, GroupId: id})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "join group failed when creating group:%s", err.Error())
	}
	return &pb.CreateGroupResp{Id: id}, nil
}
