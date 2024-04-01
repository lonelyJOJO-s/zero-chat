package userservicelogic

import (
	"context"
	"fmt"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/app/user/model"
	"zero-chat/common/tool"
	"zero-chat/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrUserNoExistsError = xerr.NewErrMsg("user not exsit")
var ErrUsernamePwdError = xerr.NewErrMsg("password error")
var ErrNoValidCaptcha = xerr.NewErrMsg("no valid captcha found in redis")
var ErrCapchaNotMatch = xerr.NewErrMsg("capcha not matched")

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// todo: add your logic here and delete this line
	var userId int64
	var err error
	switch in.Type { // 1 for pwd login 2 for email login
	case 1:
		userId, err = l.LoginByPwd(in.Account, in.Password)
		if err != nil {
			return nil, err
		}
	case 2:
		userId, err = l.LoginByEmail(in.Account, in.Password)
		if err != nil {
			return nil, err
		}
	}
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userId)
	}
	return &pb.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

func (l *LoginLogic) LoginByPwd(account, pwd string) (int64, error) {
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, account)
	if err != nil && err != model.ErrNotFound {
		return 0, errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), err.Error())
	}
	if err == model.ErrNotFound {
		user, err = l.svcCtx.UserModel.FindOneByPhone(l.ctx, account)
	}
	if err != nil && err != model.ErrNotFound {
		return 0, errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), err.Error())
	}
	if user == nil {
		return 0, errors.Wrapf(ErrUserNoExistsError, "account:%s", account)
	}
	if !(tool.Md5ByString(pwd) == user.Password) {
		return 0, errors.Wrap(ErrUsernamePwdError, "password error")
	}
	return user.Id, nil
}

func (l *LoginLogic) LoginByEmail(email string, captcha string) (int64, error) {
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, email)
	if err != nil && err != model.ErrNotFound {
		return 0, errors.Wrap(xerr.NewErrCode(xerr.DB_ERROR), err.Error())
	}
	if user == nil {
		return 0, errors.Wrapf(ErrUserNoExistsError, "email:%s", email)
	}
	// check for the vertify code
	redisKey := fmt.Sprintf("%s%s", RedisEmailLoginPrefix, email)
	val, err := l.svcCtx.RedisClient.GetCtx(l.ctx, redisKey)
	if err != nil {
		return 0, xerr.NewErrCode(xerr.REDIS_ERROR)
	}
	if val == "" {
		return 0, ErrNoValidCaptcha
	}
	if val != captcha {
		return 0, ErrCapchaNotMatch
	}
	l.svcCtx.RedisClient.DelCtx(l.ctx, redisKey)
	return user.Id, nil
}
