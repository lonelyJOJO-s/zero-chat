// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package friendservice

import (
	"context"

	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddFriendsReq          = pb.AddFriendsReq
	AddFriendsResp         = pb.AddFriendsResp
	CreateGroupReq         = pb.CreateGroupReq
	CreateGroupResp        = pb.CreateGroupResp
	DelFriendsReq          = pb.DelFriendsReq
	DelFriendsResp         = pb.DelFriendsResp
	DelUserInfoReq         = pb.DelUserInfoReq
	DelUserInfoResp        = pb.DelUserInfoResp
	DismissGroupReq        = pb.DismissGroupReq
	DismissGroupResp       = pb.DismissGroupResp
	GenerateTokenReq       = pb.GenerateTokenReq
	GenerateTokenResp      = pb.GenerateTokenResp
	GetCaptchaReq          = pb.GetCaptchaReq
	GetCaptchaResp         = pb.GetCaptchaResp
	GetFriendsReq          = pb.GetFriendsReq
	GetFriendsResp         = pb.GetFriendsResp
	GetGroupInfoReq        = pb.GetGroupInfoReq
	GetGroupInfoResp       = pb.GetGroupInfoResp
	GetJoinedGroupIdsReq   = pb.GetJoinedGroupIdsReq
	GetJoinedGroupIdsResp  = pb.GetJoinedGroupIdsResp
	GetManagedGroupIdsReq  = pb.GetManagedGroupIdsReq
	GetManagedGroupIdsResp = pb.GetManagedGroupIdsResp
	GetMemberIdsReq        = pb.GetMemberIdsReq
	GetMemberIdsResp       = pb.GetMemberIdsResp
	GetUserInfoReq         = pb.GetUserInfoReq
	GetUserInfoResp        = pb.GetUserInfoResp
	Group                  = pb.Group
	JoinGroupReq           = pb.JoinGroupReq
	JoinGroupResp          = pb.JoinGroupResp
	LoginReq               = pb.LoginReq
	LoginResp              = pb.LoginResp
	QuitGroupReq           = pb.QuitGroupReq
	QuitGroupResp          = pb.QuitGroupResp
	RegisterReq            = pb.RegisterReq
	RegisterResp           = pb.RegisterResp
	SearchFriendFuzzyReq   = pb.SearchFriendFuzzyReq
	SearchFriendFuzzyResp  = pb.SearchFriendFuzzyResp
	SearchGroupReq         = pb.SearchGroupReq
	SearchGroupResp        = pb.SearchGroupResp
	SearchUserFuzzyReq     = pb.SearchUserFuzzyReq
	SearchUserFuzzyResp    = pb.SearchUserFuzzyResp
	UpdateGroupReq         = pb.UpdateGroupReq
	UpdateGroupResp        = pb.UpdateGroupResp
	UpdateUserInfoReq      = pb.UpdateUserInfoReq
	UpdateUserInfoResp     = pb.UpdateUserInfoResp
	User                   = pb.User
	UserWithPwd            = pb.UserWithPwd

	FriendService interface {
		// friend basic
		GetFriends(ctx context.Context, in *GetFriendsReq, opts ...grpc.CallOption) (*GetFriendsResp, error)
		AddFriends(ctx context.Context, in *AddFriendsReq, opts ...grpc.CallOption) (*AddFriendsResp, error)
		DelFrineds(ctx context.Context, in *DelFriendsReq, opts ...grpc.CallOption) (*DelFriendsResp, error)
		SearchFriendFuzzy(ctx context.Context, in *SearchFriendFuzzyReq, opts ...grpc.CallOption) (*SearchFriendFuzzyResp, error)
	}

	defaultFriendService struct {
		cli zrpc.Client
	}
)

func NewFriendService(cli zrpc.Client) FriendService {
	return &defaultFriendService{
		cli: cli,
	}
}

// friend basic
func (m *defaultFriendService) GetFriends(ctx context.Context, in *GetFriendsReq, opts ...grpc.CallOption) (*GetFriendsResp, error) {
	client := pb.NewFriendServiceClient(m.cli.Conn())
	return client.GetFriends(ctx, in, opts...)
}

func (m *defaultFriendService) AddFriends(ctx context.Context, in *AddFriendsReq, opts ...grpc.CallOption) (*AddFriendsResp, error) {
	client := pb.NewFriendServiceClient(m.cli.Conn())
	return client.AddFriends(ctx, in, opts...)
}

func (m *defaultFriendService) DelFrineds(ctx context.Context, in *DelFriendsReq, opts ...grpc.CallOption) (*DelFriendsResp, error) {
	client := pb.NewFriendServiceClient(m.cli.Conn())
	return client.DelFrineds(ctx, in, opts...)
}

func (m *defaultFriendService) SearchFriendFuzzy(ctx context.Context, in *SearchFriendFuzzyReq, opts ...grpc.CallOption) (*SearchFriendFuzzyResp, error) {
	client := pb.NewFriendServiceClient(m.cli.Conn())
	return client.SearchFriendFuzzy(ctx, in, opts...)
}
