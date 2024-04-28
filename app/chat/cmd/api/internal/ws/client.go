// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ws

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"zero-chat/app/chat/cmd/api/internal/svc"
	"zero-chat/common/constant"
	"zero-chat/common/oss"
	"zero-chat/common/protocol"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
)

// 不同于一般的websocket 后端例子会把ping/pong逻辑写在后端，这里把ping/pong逻辑写在了前端

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

func (c *Client) ReadPump(svcCtx *svc.ServiceContext) {
	defer func() {
		c.Server.UnRegister <- c
		c.Conn.Close()
	}()
	for {
		c.Conn.PongHandler()
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			WsServer.UnRegister <- c
			c.Conn.Close()
			break
		}
		msg := &protocol.Message{}
		proto.Unmarshal(message, msg)
		if msg.Type == constant.HEART_BEAT {
			pong := &protocol.Message{
				Content: constant.PONG,
				Type:    constant.HEART_BEAT,
			}
			pongByte, err := proto.Marshal(pong)
			if nil != err {
				logx.Error(err)
			}
			c.Conn.WriteMessage(websocket.BinaryMessage, pongByte)
		} else {
			// send kafka
			if msg.File != nil {
				var session string
				switch msg.ChatType {
				case constant.SINGLE:
					session = SingChatStoreName("user_"+strconv.Itoa(int(msg.From)), "user_"+strconv.Itoa(int(msg.To)))
				case constant.GROUP:
					session = "group_" + strconv.Itoa(int(msg.To))
				default:
					session = "default"
				}
				msg.FileBack, err = Upload2Oss(msg.File, session, msg.Content)
				if err != nil {
					logx.Errorf("upload file to oss error:%s", err.Error())
					continue
				}
				msg.File = []byte{}
			}
			msgSend, err := proto.Marshal(msg)
			if err != nil {
				logx.Errorf("unmarshal error with proto:%s", err.Error())
				continue
			}
			err = svcCtx.KqPusherClient.Push(string(msgSend))
			if err != nil {
				logx.Error(err)
			}
		}
	}
}

// writePump pumps messages from the Hub to the websocket Connection.
//
// A goroutine running writePump is started for each Connection. The
// application ensures that there is at most one writer to a Connection by
// executing all writes from this goroutine.
func (c *Client) WritePump(svcCtx *svc.ServiceContext) {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		c.Conn.WriteMessage(websocket.BinaryMessage, message)
	}
}

func Upload2Oss(file []byte, session string, fileName string) (url string, err error) {
	// 1. upload to oss

	timeStamp := strconv.Itoa(int(time.Now().Unix()))
	path := fmt.Sprintf("file/%s/%s/%s", session, timeStamp, fileName)
	oss.Upload2Oss(file, path)
	url = os.Getenv("OSS_URI")
	url = fmt.Sprintf("%s/%s", url, path)
	return
}

func SingChatStoreName(userA, userB string) string {
	if userA > userB {
		return userB + ":" + userA
	}
	return userA + ":" + userB
}
