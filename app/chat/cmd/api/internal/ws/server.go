// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ws

import (
	"context"
	"zero-chat/app/user/cmd/rpc/client/groupservice"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/constant"
	"zero-chat/common/protocol"

	"github.com/gogo/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
)

var WsServer *Server

func init() {
	WsServer = NewServer()
}

// Hub maintains the set of active Clients and Broadcasts messages to the
// Clients.
type Server struct {
	// Registered Clients.
	Clients map[*Client]bool

	// Inbound messages from the Clients.
	Broadcast chan []byte

	// Register requests from the Clients.
	Register chan *Client

	// UnRegister requests from Clients.
	UnRegister chan *Client
}

func NewServer() *Server {
	return &Server{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Server) Run(groupService groupservice.GroupService) {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.UnRegister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			// boardcast to all Clients
			msg := &protocol.Message{}
			proto.Unmarshal(message, msg)
			if msg.ChatType == constant.SINGLE {
				for client := range h.Clients {
					if client.ClientId == msg.To {
						select {
						case client.Send <- message:
						// user off-line, quit
						default:
							close(client.Send)
							delete(h.Clients, client)
						}
					}
				}
			} else if msg.ChatType == constant.GROUP { // group chat
				members := map[int64]bool{}
				ids, err := groupService.GetMemberIds(context.Background(), &pb.GetMemberIdsReq{GroupId: msg.To})
				if err != nil {
					logx.Error(err)
					break
				}
				for _, id := range ids.Ids {
					members[id] = true
				}
				for client := range h.Clients {
					if _, ok := members[client.ClientId]; ok {
						select {
						case client.Send <- message:
						default:
							close(client.Send)
							delete(h.Clients, client)
						}
					}
				}
			}
		}
	}
}
