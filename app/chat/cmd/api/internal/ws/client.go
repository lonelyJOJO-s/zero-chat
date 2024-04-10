// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ws

import (
	"time"
	"zero-chat/common/constant"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// Client is a middleman between the websocket Connection and the Hub.
type Client struct {
	ClientId int64

	Server *Server

	// The websocket Connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte
}

// readPump pumps messages from the websocket Connection to the Hub.
//
// The application runs readPump in a per-Connection goroutine. The application
// ensures that there is at most one reader on a Connection by executing all
// reads from this goroutine.

func (c *Client) ReadPump() {
	defer func() {
		c.Server.UnRegister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			WsServer.UnRegister <- c
			c.Conn.Close()
			break
		}
		msg := &Message{}
		proto.Unmarshal(message, msg)
		if msg.Type == constant.HEART_BEAT {
			pong := &Message{
				Content: constant.PONG,
				Type:    constant.HEART_BEAT,
			}
			pongByte, err := proto.Marshal(pong)
			if nil != err {
				logx.Error(err)
			}
			c.Conn.WriteMessage(websocket.BinaryMessage, pongByte)
		} else {
			// TODO: 可以使用kafka作为消息队列中间件处理数据
			WsServer.Broadcast <- message
		}
	}
}

// writePump pumps messages from the Hub to the websocket Connection.
//
// A goroutine running writePump is started for each Connection. The
// application ensures that there is at most one writer to a Connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			// 数据回传客户端
			c.Conn.WriteMessage(websocket.BinaryMessage, message)

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
