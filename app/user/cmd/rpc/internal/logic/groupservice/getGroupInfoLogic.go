package groupservicelogic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"zero-chat/app/user/cmd/rpc/internal/svc"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/tool"
	"zero-chat/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupInfoLogic {
	return &GetGroupInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupInfoLogic) GetGroupInfo(in *pb.GetGroupInfoReq) (*pb.GetGroupInfoResp, error) {
	group, err := l.svcCtx.GroupModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find group error:%s", err.Error())
	}
	var groupPb pb.Group
	copier.Copy(&groupPb, group)
	groupPb.CreateAt = tool.Time2TimeStamp(group.CreatedAt)
	return &pb.GetGroupInfoResp{Group: &groupPb}, nil
}
