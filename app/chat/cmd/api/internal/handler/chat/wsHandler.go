package chat

import (
	"net/http"
	"zero-chat/app/chat/cmd/api/internal/logic/chat"
	"zero-chat/app/chat/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

func WsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := chat.NewWsLogic(r.Context(), svcCtx)
		err := l.Ws(w, r)
		if err != nil {
			logx.Errorf("websocket conn error:%s", err.Error())
		}
	}
}
