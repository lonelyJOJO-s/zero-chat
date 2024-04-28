package group

import (
	"context"

	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMembersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMembersLogic {
	return &GetMembersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMembersLogic) GetMembers(req *types.GetMembersReq) (resp *types.GetMembersResp, err error) {

	// todo:判断是否在群里
	membersResp, err := l.svcCtx.GroupServiceRpc.GetMemberIds(l.ctx, &pb.GetMemberIdsReq{GroupId: req.Id})
	if err != nil {
		return nil, errors.Wrapf(err, "get members error with:%s", err.Error())
	}
	if len(membersResp.Ids) == 0 {
		return
	}
	usersResp, err := l.svcCtx.UserServiceRpc.GetUsersInfo(l.ctx, &pb.GetUsersInfoReq{Ids: membersResp.Ids})
	if err != nil {
		return nil, errors.Wrapf(err, "get members info error with:%s", err.Error())
	}
	resp = new(types.GetMembersResp)
	for _, user := range usersResp.Users {
		var u types.UserBasic
		copier.Copy(&u, user)
		resp.Users = append(resp.Users, u)
	}
	return
}
