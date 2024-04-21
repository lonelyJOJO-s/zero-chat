package friend

import (
	"context"

	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/ctxdata"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FriendDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendDelLogic {
	return &FriendDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendDelLogic) FriendDel(req *types.FriendIdReq) (resp *types.Null, err error) {
	id := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.FriendServiceRpc.DelFrineds(l.ctx, &pb.DelFriendsReq{Id: id, FriendId: req.Id})
	if err != nil {
		return nil, errors.Wrapf(err, "del friend rpc error with:%s", err.Error())
	}
	return
}
