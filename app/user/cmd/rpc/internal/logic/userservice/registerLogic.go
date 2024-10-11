package userservicelogic

import (
	"context"
	"fmt"
	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/app/user/model"
	"zero-chat/common/tool"
	"zero-chat/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrGenerateTokenError = xerr.NewErrMsg("generate token failed")
var ErrUserAlreadyRegisterError = xerr.NewErrMsg("user has been registered")

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.UserInfo.Username)
	// db error
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "username:%s,err:%v", in.UserInfo.Username, err)
	}
	// user exsit
	if user != nil {
		return nil, errors.Wrapf(ErrUserAlreadyRegisterError, "Register user exists username:%s,err:%v", in.UserInfo.Username, err)
	}
	var userId int64
	user = new(model.Users)
	copier.Copy(user, in.UserInfo)
	user.Avatar = in.UserInfo.Avatar
	if len(in.UserInfo.Password) > 0 {
		user.Password = tool.Md5ByString(in.UserInfo.Password)
	}
	insertResult, err := l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user Insert err:%v,user:%+v", err, user)
	}
	lastId, err := insertResult.LastInsertId()
	fmt.Println(lastId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Register db user insertResult.LastInsertId err:%v,user:%+v", err, user)
	}
	userId = lastId

	//2、Generate the token, so that the service doesn't call rpc internally
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userId)
	}
	l.Logger.Info("%s has been registerd successfully.", in.UserInfo.Username)
	return &pb.RegisterResp{
		Id:           userId,
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
