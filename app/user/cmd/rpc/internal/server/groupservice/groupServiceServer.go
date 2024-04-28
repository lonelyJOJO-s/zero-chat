// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package server

import (
	"context"

	"zero-chat/app/user/cmd/rpc/internal/logic/groupservice"
	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
)

type GroupServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedGroupServiceServer
}

func NewGroupServiceServer(svcCtx *svc.ServiceContext) *GroupServiceServer {
	return &GroupServiceServer{
		svcCtx: svcCtx,
	}
}

// group
func (s *GroupServiceServer) CreateGroup(ctx context.Context, in *pb.CreateGroupReq) (*pb.CreateGroupResp, error) {
	l := groupservicelogic.NewCreateGroupLogic(ctx, s.svcCtx)
	return l.CreateGroup(in)
}

func (s *GroupServiceServer) JoinGroup(ctx context.Context, in *pb.JoinGroupReq) (*pb.JoinGroupResp, error) {
	l := groupservicelogic.NewJoinGroupLogic(ctx, s.svcCtx)
	return l.JoinGroup(in)
}

func (s *GroupServiceServer) QuitGroup(ctx context.Context, in *pb.QuitGroupReq) (*pb.QuitGroupResp, error) {
	l := groupservicelogic.NewQuitGroupLogic(ctx, s.svcCtx)
	return l.QuitGroup(in)
}

func (s *GroupServiceServer) DismissGroup(ctx context.Context, in *pb.DismissGroupReq) (*pb.DismissGroupResp, error) {
	l := groupservicelogic.NewDismissGroupLogic(ctx, s.svcCtx)
	return l.DismissGroup(in)
}

func (s *GroupServiceServer) UpdateGroupInfo(ctx context.Context, in *pb.UpdateGroupReq) (*pb.UpdateGroupResp, error) {
	l := groupservicelogic.NewUpdateGroupInfoLogic(ctx, s.svcCtx)
	return l.UpdateGroupInfo(in)
}

func (s *GroupServiceServer) GetGroupInfo(ctx context.Context, in *pb.GetGroupInfoReq) (*pb.GetGroupInfoResp, error) {
	l := groupservicelogic.NewGetGroupInfoLogic(ctx, s.svcCtx)
	return l.GetGroupInfo(in)
}

func (s *GroupServiceServer) SearchGroup(ctx context.Context, in *pb.SearchGroupReq) (*pb.SearchGroupResp, error) {
	l := groupservicelogic.NewSearchGroupLogic(ctx, s.svcCtx)
	return l.SearchGroup(in)
}

func (s *GroupServiceServer) GetMemberIds(ctx context.Context, in *pb.GetMemberIdsReq) (*pb.GetMemberIdsResp, error) {
	l := groupservicelogic.NewGetMemberIdsLogic(ctx, s.svcCtx)
	return l.GetMemberIds(in)
}

func (s *GroupServiceServer) GetManagedGroupIds(ctx context.Context, in *pb.GetManagedGroupIdsReq) (*pb.GetManagedGroupIdsResp, error) {
	l := groupservicelogic.NewGetManagedGroupIdsLogic(ctx, s.svcCtx)
	return l.GetManagedGroupIds(in)
}

func (s *GroupServiceServer) GetJoinedGroupIds(ctx context.Context, in *pb.GetJoinedGroupIdsReq) (*pb.GetJoinedGroupIdsResp, error) {
	l := groupservicelogic.NewGetJoinedGroupIdsLogic(ctx, s.svcCtx)
	return l.GetJoinedGroupIds(in)
}

func (s *GroupServiceServer) GetJoinedGroups(ctx context.Context, in *pb.GetJoinedGroupsReq) (*pb.GetJoinedGroupsResp, error) {
	l := groupservicelogic.NewGetJoinedGroupsLogic(ctx, s.svcCtx)
	return l.GetJoinedGroups(in)
}

func (s *GroupServiceServer) SearchAllGroup(ctx context.Context, in *pb.SearchAllGroupReq) (*pb.SearchAllGroupResp, error) {
	l := groupservicelogic.NewSearchAllGroupLogic(ctx, s.svcCtx)
	return l.SearchAllGroup(in)
}
