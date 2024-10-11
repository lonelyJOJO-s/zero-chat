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

type SearchAllGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchAllGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchAllGroupLogic {
	return &SearchAllGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchAllGroupLogic) SearchAllGroup(req *types.SearchAllGroupReq) (resp *types.SearchAllGroupResp, err error) {
	// todo: add your logic here and delete this line
	groupResp, err := l.svcCtx.GroupServiceRpc.SearchAllGroup(l.ctx, &pb.SearchAllGroupReq{Keyword: req.Keyword})
	if err != nil {
		return nil, errors.Wrapf(err, "search all group error with:%s", err.Error())
	}
	joinedGroupResp, err := l.svcCtx.GroupServiceRpc.GetJoinedGroupIds(l.ctx, &pb.GetJoinedGroupIdsReq{UserId: ctxdata.GetUidFromCtx(l.ctx)})
	if err != nil {
		return nil, errors.Wrapf(err, "search joined group error with:%s", err.Error())
	}
	joinedIdsMap := map[int64]bool{}
	for _, jid := range joinedGroupResp.Ids {
		joinedIdsMap[jid] = true
	}
	resp = new(types.SearchAllGroupResp)
	for _, group := range groupResp.Group {
		if !joinedIdsMap[group.Id] {
			var g types.GroupWithId
			copier.Copy(&g, group)
			g.CreatorId = group.OwnerId
			resp.Groups = append(resp.Groups, g)
		}
	}
	return
}
