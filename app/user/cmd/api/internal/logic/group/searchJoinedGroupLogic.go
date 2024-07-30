package group

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

type SearchJoinedGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchJoinedGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchJoinedGroupLogic {
	return &SearchJoinedGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchJoinedGroupLogic) SearchJoinedGroup(req *types.SearchJoinedGroupReq) (resp *types.SearchJoinedGroupResp, err error) {
	id := ctxdata.GetUidFromCtx(l.ctx)
	rpcResp, err := l.svcCtx.GroupServiceRpc.SearchGroup(l.ctx, &pb.SearchGroupReq{UserId: id, Keyword: req.Keyword})
	if err != nil {
		return nil, errors.Wrapf(err, "search group error with:%s", err.Error())
	}
	resp = new(types.SearchJoinedGroupResp)
	for _, group := range rpcResp.Group {
		var g types.GroupWithId
		copier.Copy(&g, group)
		g.CreatorId = group.OwnerId
		resp.Groups = append(resp.Groups, g)
	}
	return
}
