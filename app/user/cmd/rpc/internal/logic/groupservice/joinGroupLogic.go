package groupservicelogic

import (
	"context"
	"time"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/app/user/model"
	"zero-chat/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type JoinGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupLogic {
	return &JoinGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinGroupLogic) JoinGroup(in *pb.JoinGroupReq) (*pb.JoinGroupResp, error) {
	// judge wheather in db already
	userGroup, err := l.svcCtx.UserGroup.FindOneByUserIdGroupId(l.ctx, in.UserId, in.GroupId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find usergroup error:%s", err.Error())
	}
	if userGroup != nil {
		return nil, xerr.NewErrCodeMsg(xerr.INSERT_ALREADY_EXSIT, xerr.MapErrMsg(xerr.INSERT_ALREADY_EXSIT))
	}
	_, err = l.svcCtx.UserGroup.Insert(l.ctx, &model.UserGroup{GroupId: in.GroupId, UserId: in.UserId, CreatedAt: time.Now()})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "join group error:%s", err.Error())
	}
	return &pb.JoinGroupResp{GroupId: in.GroupId}, nil
}
