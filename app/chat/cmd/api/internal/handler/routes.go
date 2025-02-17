// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	chat "zero-chat/app/chat/cmd/api/internal/handler/chat"
	"zero-chat/app/chat/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Jwt},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/ws",
					Handler: chat.WsHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/chat/api/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/history-message",
				Handler: chat.GetHistoryMessageHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/history-message/search",
				Handler: chat.SearchHistoryMessageHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/chat/api/v1"),
		rest.WithTimeout(3000*time.Millisecond),
	)
}
