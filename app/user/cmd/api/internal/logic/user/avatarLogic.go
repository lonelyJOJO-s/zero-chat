package user

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

	"github.com/zeromicro/go-zero/core/logx"
)

type AvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AvatarLogic {
	return &AvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AvatarLogic) Avatar(req *types.Null, r *http.Request) (resp string, err error) {
	// todo: add your logic here and delete this line
	_, file, err := r.FormFile("file")
	if err != nil {
		return
	}
	resp, err = l.Upload2Oss(file, int(ctxdata.GetUidFromCtx(l.ctx)))
	if err != nil {
		return
	}
	logx.Infof("upload file successfully. url:%s", resp)
	return
}

func (l *AvatarLogic) Upload2Oss(file *multipart.FileHeader, userId int) (url string, err error) {
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
	go oss.Upload2Oss(fileBytes, path)
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
