package friendservicelogic

import (
	"context"
	"time"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLastTimeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLastTimeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLastTimeLogic {
	return &UpdateLastTimeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLastTimeLogic) UpdateLastTime(in *pb.UpdateLastTimeReq) (*pb.UpdateLastTimeResp, error) {
	from, to := in.From, in.To
	uf1, err := l.svcCtx.UserFriend.FindOneByUserIdFriendId(l.ctx, from, to)
	if err != nil {
		return nil, errors.Wrapf(err, "find userfriend relation error")
	}
	uf2, err := l.svcCtx.UserFriend.FindOneByUserIdFriendId(l.ctx, to, from)
	if err != nil {
		return nil, errors.Wrapf(err, "find userfriend relation error")
	}
	uf1.LastMessageTime = time.Now()
	uf2.LastMessageTime = uf1.LastMessageTime
	l.svcCtx.UserFriend.Update(l.ctx, uf1)
	l.svcCtx.UserFriend.Update(l.ctx, uf2)
	return &pb.UpdateLastTimeResp{}, nil
}
