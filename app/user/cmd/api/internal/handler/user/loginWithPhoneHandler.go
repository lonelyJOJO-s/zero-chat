package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"zero-chat/app/user/cmd/api/internal/logic/user"
	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/common/result"
)

func LoginWithPhoneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PhoneLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// validator
		validate := validator.New()
		if err := validate.StructCtx(r.Context(), req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewLoginWithPhoneLogic(r.Context(), svcCtx)
		resp, err := l.LoginWithPhone(&req)
		// uniform return
		result.HttpResult(r, w, resp, err)
	}
}
