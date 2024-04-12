package friendservicelogic

import (
	"context"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUuidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUuidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUuidLogic {
	return &GetUuidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUuidLogic) GetUuid(in *pb.GetUuidReq) (*pb.GetUuidResp, error) {
	// todo: add your logic here and delete this line
	userFriend, err := l.svcCtx.UserFriend.FindOneByUserIdFriendId(l.ctx, in.UserId, in.FriendId)
	if err != nil {
		return nil, errors.Wrapf(err, "get uuid error:%s", err.Error())
	}
	return &pb.GetUuidResp{Uuid: userFriend.Uuid}, nil
}
