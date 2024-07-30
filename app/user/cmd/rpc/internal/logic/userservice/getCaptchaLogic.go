package userservicelogic

import (
	"context"
	"fmt"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/email"
	"zero-chat/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrSendEmailError = xerr.NewErrMsg("email send fail")
var ErrRepeatedSend = xerr.NewErrMsg("email send too frequetially")
var RedisEmailLoginPrefix = "cache:login:email:"

type GetCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCaptchaLogic) GetCaptcha(in *pb.GetCaptchaReq) (*pb.GetCaptchaResp, error) {
	redisKey := fmt.Sprintf("%s%s", RedisEmailLoginPrefix, in.Email)
	// TODO: 这一块限流放在api层（rpc一般不会对外访问）
	// val, err := l.svcCtx.RedisClient.GetCtx(l.ctx, redisKey)
	// if err != nil {
	// 	return nil, errors.Wrapf(xerr.NewErrCode(xerr.REDIS_ERROR), "get captcha use redis fail:%s", err.Error())
	// }
	// if val != "" {
	// 	return nil, ErrRepeatedSend
	// }
	code, err := email.SendEmail([]string{in.Email}, email.LoginType)
	if err != nil {
		return nil, errors.Wrapf(ErrSendEmailError, "something went wrong in the email block: %s", err.Error())
	}
	// store into redis
	l.svcCtx.RedisClient.SetexCtx(l.ctx, redisKey, code, 60*3)
	return &pb.GetCaptchaResp{
		Code: code,
	}, nil
}
