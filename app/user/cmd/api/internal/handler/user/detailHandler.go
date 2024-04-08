package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zero-chat/app/user/cmd/api/internal/logic/user"
	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/common/result"
	"zero-chat/common/xerr"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.HttpResult(r, w, nil, errors.Wrapf(xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR), "params errors with:%s", err.Error()))
			return
		}
		// validator
		validate := validator.New()
		if err := validate.StructCtx(r.Context(), req); err != nil {
			result.HttpResult(r, w, nil, errors.Wrapf(xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR), "params errors with:%s", err.Error()))
			return
		}

		l := user.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail(&req)
		// uniform return
		result.HttpResult(r, w, resp, err)
	}
}
