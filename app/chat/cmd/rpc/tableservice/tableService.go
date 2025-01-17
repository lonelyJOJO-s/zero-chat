// Code generated by goctl. DO NOT EDIT.
// Source: chat.proto

package tableservice

import (
	"context"

	"zero-chat/app/chat/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetHistoryMessageReq     = pb.GetHistoryMessageReq
	GetHistoryMessageResp    = pb.GetHistoryMessageResp
	GetSyncMessageReq        = pb.GetSyncMessageReq
	GetSyncMessageResp       = pb.GetSyncMessageResp
	Message                  = pb.Message
	MessageWithSeq           = pb.MessageWithSeq
	Null                     = pb.Null
	SearchHistoryMessageReq  = pb.SearchHistoryMessageReq
	SearchHistoryMessageResp = pb.SearchHistoryMessageResp
	SendReq                  = pb.SendReq
	SendResp                 = pb.SendResp
	StoreAddItemReq          = pb.StoreAddItemReq
	StoreAddItemResp         = pb.StoreAddItemResp
	StoreTableItem           = pb.StoreTableItem
	SyncAddItemReq           = pb.SyncAddItemReq
	SyncAddItemResp          = pb.SyncAddItemResp
	SyncTableItem            = pb.SyncTableItem

	TableService interface {
		// has been depricated
		StoreAddItem(ctx context.Context, in *StoreAddItemReq, opts ...grpc.CallOption) (*StoreAddItemResp, error)
		SyncAddItem(ctx context.Context, in *SyncAddItemReq, opts ...grpc.CallOption) (*SyncAddItemResp, error)
		Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*SendResp, error)
		GetSyncMessage(ctx context.Context, in *GetSyncMessageReq, opts ...grpc.CallOption) (*GetSyncMessageResp, error)
		GetHistoryMessage(ctx context.Context, in *GetHistoryMessageReq, opts ...grpc.CallOption) (*GetHistoryMessageResp, error)
		SearchHistoryMEssage(ctx context.Context, in *SearchHistoryMessageReq, opts ...grpc.CallOption) (*SearchHistoryMessageResp, error)
	}

	defaultTableService struct {
		cli zrpc.Client
	}
)

func NewTableService(cli zrpc.Client) TableService {
	return &defaultTableService{
		cli: cli,
	}
}

// has been depricated
func (m *defaultTableService) StoreAddItem(ctx context.Context, in *StoreAddItemReq, opts ...grpc.CallOption) (*StoreAddItemResp, error) {
	client := pb.NewTableServiceClient(m.cli.Conn())
	return client.StoreAddItem(ctx, in, opts...)
}

func (m *defaultTableService) SyncAddItem(ctx context.Context, in *SyncAddItemReq, opts ...grpc.CallOption) (*SyncAddItemResp, error) {
	client := pb.NewTableServiceClient(m.cli.Conn())
	return client.SyncAddItem(ctx, in, opts...)
}

func (m *defaultTableService) Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*SendResp, error) {
	client := pb.NewTableServiceClient(m.cli.Conn())
	return client.Send(ctx, in, opts...)
}

func (m *defaultTableService) GetSyncMessage(ctx context.Context, in *GetSyncMessageReq, opts ...grpc.CallOption) (*GetSyncMessageResp, error) {
	client := pb.NewTableServiceClient(m.cli.Conn())
	return client.GetSyncMessage(ctx, in, opts...)
}

func (m *defaultTableService) GetHistoryMessage(ctx context.Context, in *GetHistoryMessageReq, opts ...grpc.CallOption) (*GetHistoryMessageResp, error) {
	client := pb.NewTableServiceClient(m.cli.Conn())
	return client.GetHistoryMessage(ctx, in, opts...)
}

func (m *defaultTableService) SearchHistoryMEssage(ctx context.Context, in *SearchHistoryMessageReq, opts ...grpc.CallOption) (*SearchHistoryMessageResp, error) {
	client := pb.NewTableServiceClient(m.cli.Conn())
	return client.SearchHistoryMEssage(ctx, in, opts...)
}
