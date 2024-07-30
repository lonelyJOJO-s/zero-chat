package friend

import (
	"net/http"
	"zero-chat/app/user/cmd/api/internal/logic/friend"
	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/common/result"
	"zero-chat/common/xerr"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func SearchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendSearchReq
		val := r.URL.Query()
		keyword := val.Get("keyword")
		req.Keyword = keyword
		// validator
		validate := validator.New()
		if err := validate.StructCtx(r.Context(), req); err != nil {
			result.HttpResult(r, w, nil, errors.Wrapf(xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR), "validate errors with:%s", err.Error()))
			return
		}

		l := friend.NewSearchLogic(r.Context(), svcCtx)
		resp, err := l.Search(&req)
		// uniform return
		result.HttpResult(r, w, resp, err)
	}
}
