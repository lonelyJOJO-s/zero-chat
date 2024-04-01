// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package userservice

import (
	"context"

	"zero-chat/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddFriendsReq         = pb.AddFriendsReq
	AddFriendsResp        = pb.AddFriendsResp
	CreateGroupReq        = pb.CreateGroupReq
	CreateGroupResp       = pb.CreateGroupResp
	DelFriendsReq         = pb.DelFriendsReq
	DelFriendsResp        = pb.DelFriendsResp
	DelUserInfoReq        = pb.DelUserInfoReq
	DelUserInfoResp       = pb.DelUserInfoResp
	DismissGroupReq       = pb.DismissGroupReq
	DismissGroupResp      = pb.DismissGroupResp
	GenerateTokenReq      = pb.GenerateTokenReq
	GenerateTokenResp     = pb.GenerateTokenResp
	GetCaptchaReq         = pb.GetCaptchaReq
	GetCaptchaResp        = pb.GetCaptchaResp
	GetFriendsReq         = pb.GetFriendsReq
	GetFriendsResp        = pb.GetFriendsResp
	GetGroupInfoReq       = pb.GetGroupInfoReq
	GetGroupInfoResp      = pb.GetGroupInfoResp
	GetUserInfoReq        = pb.GetUserInfoReq
	GetUserInfoResp       = pb.GetUserInfoResp
	Group                 = pb.Group
	JoinGroupReq          = pb.JoinGroupReq
	JoinGroupResp         = pb.JoinGroupResp
	LoginReq              = pb.LoginReq
	LoginResp             = pb.LoginResp
	QuitGroupReq          = pb.QuitGroupReq
	QuitGroupResp         = pb.QuitGroupResp
	RegisterReq           = pb.RegisterReq
	RegisterResp          = pb.RegisterResp
	SearchFriendFuzzyReq  = pb.SearchFriendFuzzyReq
	SearchFriendFuzzyResp = pb.SearchFriendFuzzyResp
	SearchGroupReq        = pb.SearchGroupReq
	SearchGroupResp       = pb.SearchGroupResp
	SearchUserFuzzyReq    = pb.SearchUserFuzzyReq
	SearchUserFuzzyResp   = pb.SearchUserFuzzyResp
	UpdateGroupReq        = pb.UpdateGroupReq
	UpdateGroupResp       = pb.UpdateGroupResp
	UpdateUserInfoReq     = pb.UpdateUserInfoReq
	UpdateUserInfoResp    = pb.UpdateUserInfoResp
	User                  = pb.User
	UserWithPwd           = pb.UserWithPwd

	UserService interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
		// user basic
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
		SoftDelUser(ctx context.Context, in *DelUserInfoReq, opts ...grpc.CallOption) (*DelUserInfoResp, error)
		UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UpdateUserInfoResp, error)
		SearchUserFuzzy(ctx context.Context, in *SearchUserFuzzyReq, opts ...grpc.CallOption) (*SearchUserFuzzyResp, error)
		GetCaptcha(ctx context.Context, in *GetCaptchaReq, opts ...grpc.CallOption) (*GetCaptchaResp, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUserService) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUserService) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

// user basic
func (m *defaultUserService) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUserService) SoftDelUser(ctx context.Context, in *DelUserInfoReq, opts ...grpc.CallOption) (*DelUserInfoResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.SoftDelUser(ctx, in, opts...)
}

func (m *defaultUserService) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UpdateUserInfoResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.UpdateUserInfo(ctx, in, opts...)
}

func (m *defaultUserService) SearchUserFuzzy(ctx context.Context, in *SearchUserFuzzyReq, opts ...grpc.CallOption) (*SearchUserFuzzyResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.SearchUserFuzzy(ctx, in, opts...)
}

func (m *defaultUserService) GetCaptcha(ctx context.Context, in *GetCaptchaReq, opts ...grpc.CallOption) (*GetCaptchaResp, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.GetCaptcha(ctx, in, opts...)
}