package userservicelogic

import (
	"context"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/app/user/model"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserFuzzyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserFuzzyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserFuzzyLogic {
	return &SearchUserFuzzyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserFuzzyLogic) SearchUserFuzzy(in *pb.SearchUserFuzzyReq) (*pb.SearchUserFuzzyResp, error) {
	// search user by partial phone, email, username
	r1, err := l.svcCtx.UserModel.FuzzyFindByEmail(l.ctx, in.Keyword)
	if err != nil {
		return nil, err
	}
	r2, err := l.svcCtx.UserModel.FuzzyFindByPhone(l.ctx, in.Keyword)
	if err != nil {
		return nil, err
	}
	r3, err := l.svcCtx.UserModel.FuzzyFindByUsername(l.ctx, in.Keyword)
	if err != nil {
		return nil, err
	}
	users := make([]*model.Users, 0)
	if r1 != nil {
		users = append(users, r1...)
	}

	if r2 != nil {
		users = append(users, r2...)
	}

	if r3 != nil {
		users = append(users, r3...)
	}
	var result = new(pb.SearchUserFuzzyResp)
	// TODO: slice copy error, figure out why?
	for _, user := range users {
		t := new(pb.User)
		copier.Copy(t, user)
		result.Users = append(result.Users, t)
	}
	return result, nil
}
