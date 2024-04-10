// Code generated by goctl. DO NOT EDIT.
// Source: chat.proto

package chatservice

import (
	"context"

	"zero-chat/app/chat/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Message = pb.Message
	Null    = pb.Null
	SendReq = pb.SendReq

	ChatService interface {
		Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*Null, error)
	}

	defaultChatService struct {
		cli zrpc.Client
	}
)

func NewChatService(cli zrpc.Client) ChatService {
	return &defaultChatService{
		cli: cli,
	}
}

func (m *defaultChatService) Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*Null, error) {
	client := pb.NewChatServiceClient(m.cli.Conn())
	return client.Send(ctx, in, opts...)
}
