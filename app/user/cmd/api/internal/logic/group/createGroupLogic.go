package group

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"zero-chat/app/user/cmd/api/internal/svc"
	"zero-chat/app/user/cmd/api/internal/types"
	"zero-chat/app/user/cmd/rpc/pb"
	"zero-chat/common/ctxdata"
	"zero-chat/common/oss"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupLogic) CreateGroup(req *types.GroupCreateReq, r *http.Request) (resp *types.GroupResp, err error) {
	// todo: add your logic here and delete this line
	_, file, err := r.FormFile("file")
	if err != nil {
		return
	}
	avatar, err := l.Upload2Oss(file, int(ctxdata.GetUidFromCtx(l.ctx)))
	if err != nil {
		return
	}
	pbResp, err := l.svcCtx.GroupServiceRpc.CreateGroup(l.ctx, &pb.CreateGroupReq{Keyword: &pb.Group{
		Name:    req.Name,
		Desc:    req.Desc,
		Avatar:  avatar,
		OwnerId: ctxdata.GetUidFromCtx(l.ctx),
	}})
	if err != nil {
		return nil, errors.Wrapf(err, "create group error with:%s", err.Error())
	}
	resp = &types.GroupResp{
		Group: types.GroupWithId{
			Id: pbResp.Id,
			Group: types.Group{
				Name:      req.Name,
				Desc:      req.Desc,
				Avatar:    avatar,
				CreatorId: ctxdata.GetUidFromCtx(l.ctx),
			},
		},
	}
	return
}

func (l *CreateGroupLogic) Upload2Oss(file *multipart.FileHeader, userId int) (url string, err error) {
	// 1. upload to oss
	openedFile, err := file.Open()
	if err != nil {
		return
	}
	defer openedFile.Close()

	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		return
	}
	timeStamp := strconv.Itoa(int(time.Now().Unix()))
	path := fmt.Sprintf("avatar/user/%d/%s/%s", userId, timeStamp, file.Filename)
	oss.Upload2Oss(fileBytes, path)
	url = os.Getenv("OSS_URI")
	url = fmt.Sprintf("%s/%s", url, path)
	// 2. store to local db
	_, err = l.svcCtx.UserServiceRpc.UpdateUserInfo(l.ctx, &pb.UpdateUserInfoReq{
		User: &pb.UserWithPwd{Id: int64(userId), Avatar: url},
	})
	if err != nil {
		return
	}
	return
}
