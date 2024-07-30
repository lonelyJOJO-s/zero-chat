package logic

import (
	"context"

	"zero-chat/app/chat/cmd/rpc/internal/svc"
	"zero-chat/app/chat/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreAddItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStoreAddItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreAddItemLogic {
	return &StoreAddItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// store table
func (l *StoreAddItemLogic) StoreAddItem(in *pb.StoreAddItemReq) (*pb.StoreAddItemResp, error) {
	// sequenceId := time.Now().UnixNano()
	// err := l.svcCtx.StoreTable.PutRow(
	// 	model.NewPrimaryKeys(in.Msg.TimeLineId, sequenceId),
	// 	model.NewColumns(model.WithContent(in.Msg.Content), model.WithConversation(in.Msg.Conversation),
	// 		model.WithFile(in.Msg.File), model.WithMsgType(in.Msg.MsgType), model.WithType(in.Msg.ContentType),
	// 		model.WithSender(in.Msg.Sender), model.WithSendTime(in.Msg.SendTime)),
	// )
	// if err != nil {
	// 	return nil, errors.Wrapf(err, "insert into store-table error:%s", err.Error())
	// }

	return &pb.StoreAddItemResp{TimeLineId: in.Msg.TimeLineId}, nil
}
