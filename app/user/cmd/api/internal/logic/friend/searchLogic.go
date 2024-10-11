package friend

import (
	"context"

	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.FriendSearchReq) (resp *types.FriendSearchResp, err error) {
	// todo: add your logic here and delete this line
	id := ctxdata.GetUidFromCtx(l.ctx)
	friendsResp, err := l.svcCtx.FriendServiceRpc.SearchFriendFuzzy(l.ctx, &pb.SearchFriendFuzzyReq{Id: id, Keyword: req.Keyword})
	if err != nil {
		return nil, errors.Wrapf(err, "del friend rpc error with:%s", err.Error())
	}
	resp = new(types.FriendSearchResp)
	for _, user := range friendsResp.Users {
		u := types.User{}
		copier.Copy(&u, user)
		resp.Users = append(resp.Users, u)
	}
	return
}
