package friendservicelogic

import (
	"context"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchFriendFuzzyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchFriendFuzzyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchFriendFuzzyLogic {
	return &SearchFriendFuzzyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchFriendFuzzyLogic) SearchFriendFuzzy(in *pb.SearchUserFuzzyReq) (*pb.SearchUserFuzzyResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchUserFuzzyResp{}, nil
}
