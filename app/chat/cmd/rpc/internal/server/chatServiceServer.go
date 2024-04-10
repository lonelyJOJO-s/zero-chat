// Code generated by goctl. DO NOT EDIT.
// Source: chat.proto

package server

import (
	"context"

	"zero-chat/app/chat/cmd/rpc/internal/logic"
	"zero-chat/app/chat/cmd/rpc/internal/svc"
	"zero-chat/app/chat/cmd/rpc/pb"
)

type ChatServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedChatServiceServer
}

func NewChatServiceServer(svcCtx *svc.ServiceContext) *ChatServiceServer {
	return &ChatServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *ChatServiceServer) Send(ctx context.Context, in *pb.SendReq) (*pb.Null, error) {
	l := logic.NewSendLogic(ctx, s.svcCtx)
	return l.Send(in)
}