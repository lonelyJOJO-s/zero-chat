package chat

import (
	"context"
	"net/http"

	"zero-chat/app/chat/cmd/api/internal/svc"
	"zero-chat/app/chat/cmd/api/internal/ws"
	"zero-chat/common/ctxdata"
	"zero-chat/common/xerr"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type WsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

const bufSize = 256

func NewWsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsLogic {
	return &WsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (l *WsLogic) Ws(w http.ResponseWriter, r *http.Request) error {
	// upgrade to websocket
	userId := ctxdata.GetUidFromCtx(l.ctx)
	logx.Infof("user:%d has been connected to ws", userId)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.WEBSOCKET_CONN_ERR), "conn to websocket failed:%s", err.Error())
	}
	client := &ws.Client{
		ClientId: userId,
		Server:   ws.WsServer,
		Conn:     conn,
		Send:     make(chan []byte, bufSize),
	}
	// register client
	client.Server.Register <- client
	logx.Infof("user:%d has been registered to wsServer", userId)
	// listen to websocket conn
	go client.WritePump(l.svcCtx)
	go client.ReadPump(l.svcCtx)
	return nil
}
