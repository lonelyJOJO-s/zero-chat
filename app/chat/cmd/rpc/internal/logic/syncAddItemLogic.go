package logic

import (
	"context"

	"zero-chat/app/chat/cmd/rpc/internal/svc"
	"zero-chat/app/chat/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncAddItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncAddItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncAddItemLogic {
	return &SyncAddItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// sync table
func (l *SyncAddItemLogic) SyncAddItem(in *pb.SyncAddItemReq) (*pb.SyncAddItemResp, error) {
	// todo: add your logic here and delete this line
	// sequenceId := time.Now().UnixNano()
	// err := l.svcCtx.SyncTable.PutRow(model.NewPrimaryKeys(in.Msg.TimeLineId, sequenceId),
	// 	model.NewColumns(model.WithContent(in.Msg.Content), model.WithId(strconv.Itoa(int(in.Msg.UserId))),
	// 		model.WithMsgType(in.Msg.MsgType), model.WithType(in.Msg.ContentType),
	// 		model.WithSender(in.Msg.Sender), model.WithSendTime(in.Msg.SendTime), model.WithFile(in.Msg.File)))
	// if err != nil {
	// 	return nil, errors.Wrapf(err, "insert into sync-table error:%s", err.Error())
	// }
	return &pb.SyncAddItemResp{TimeLineId: in.Msg.TimeLineId}, nil
}
