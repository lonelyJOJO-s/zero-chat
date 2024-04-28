package group

import (
	"net/http"
	"zero-chat/app/user/cmd/api/internal/logic/group"
	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/common/result"
	"zero-chat/common/tool"
	"zero-chat/common/xerr"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

func GetMembersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMembersReq
		if err := tool.ParseQuery(r, &req); err != nil {
			result.HttpResult(r, w, nil, errors.Wrapf(xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR), "params errors with:%s", err.Error()))
			return
		}
		// validator
		validate := validator.New()
		if err := validate.StructCtx(r.Context(), req); err != nil {
			result.HttpResult(r, w, nil, errors.Wrapf(xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR), "params errors with:%s", err.Error()))
			return
		}

		l := group.NewGetMembersLogic(r.Context(), svcCtx)
		resp, err := l.GetMembers(&req)
		// uniform return
		result.HttpResult(r, w, resp, err)
	}
}
