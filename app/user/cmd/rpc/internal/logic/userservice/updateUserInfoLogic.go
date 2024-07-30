package userservicelogic

import (
	"context"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/app/user/model"
	"zero-chat/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *pb.UpdateUserInfoReq) (*pb.UpdateUserInfoResp, error) {
	data, err := l.svcCtx.UserModel.FindOne(l.ctx, in.User.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(xerr.NewErrMsg("find user error"), err.Error())
	}
	if data == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "user_id:%d not found", in.User.Id)
	}
	copier.CopyWithOption(data, in.User, copier.Option{IgnoreEmpty: true})
	err = l.svcCtx.UserModel.Update(l.ctx, data)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserInfoResp{}, nil
}
