// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ws

import (
	"zero-chat/common/constant"

	"github.com/gogo/protobuf/proto"
)

var WsServer *Server

func init() {
	WsServer = NewServer()
	go WsServer.Run()
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

func (h *Server) Run() {
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
			msg := &Message{}
			proto.Unmarshal(message, msg)
			if msg.ChatType == constant.SINGLE && len(msg.To) == 1 { // solo
				for client := range h.Clients {
					if client.ClientId == msg.To[0] {
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
				for _, id := range msg.To {
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
