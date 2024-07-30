package groupservicelogic

import (
	"context"
	"strings"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchAllGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchAllGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchAllGroupLogic {
	return &SearchAllGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchAllGroupLogic) SearchAllGroup(in *pb.SearchAllGroupReq) (resp *pb.SearchAllGroupResp, err error) {
	// todo: add your logic here and delete this line
	rowBuilder := l.svcCtx.GroupModel.SelectBuilder()
	groups, err := l.svcCtx.GroupModel.FindAll(l.ctx, rowBuilder, "")
	resp = new(pb.SearchAllGroupResp)
	for _, group := range groups {
		if strings.Contains(group.Name.String, in.Keyword) {
			var pbGroup pb.Group
			copier.Copy(&pbGroup, group)
			resp.Group = append(resp.Group, &pbGroup)
		}
	}
	return
}
