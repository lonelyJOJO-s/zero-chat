package user

import (
	"net/http"
	"zero-chat/app/user/cmd/api/internal/logic/user"
	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/common/result"
)

func AvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Null

		l := user.NewAvatarLogic(r.Context(), svcCtx)
		resp, err := l.Avatar(&req, r)
		// uniform return
		result.HttpResult(r, w, resp, err)
	}
}
