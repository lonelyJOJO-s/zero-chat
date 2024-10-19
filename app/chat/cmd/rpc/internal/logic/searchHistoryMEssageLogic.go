package logic

import (
	"context"
	"strconv"

	"zero-chat/app/chat/cmd/rpc/internal/im"
	"zero-chat/app/chat/cmd/rpc/internal/svc"
	"zero-chat/app/chat/cmd/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchHistoryMEssageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchHistoryMEssageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchHistoryMEssageLogic {
	return &SearchHistoryMEssageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchHistoryMEssageLogic) SearchHistoryMEssage(in *pb.SearchHistoryMessageReq) (resp *pb.SearchHistoryMessageResp, err error) {
	var datas []im.CustomData
	var key string
	if in.GroupId > 0 {
		key = "group_" + strconv.Itoa(int(in.GroupId))
	} else {
		key = im.SingChatStoreName("user_"+strconv.Itoa(int(in.UserA)), "user_"+strconv.Itoa(int(in.UserB)))
	}
	datas, err = l.svcCtx.IM.QueryWithPrimaryKeyAndContent("im_timeline_store", "im_pk_content_index",
		key, in.KeyWord, int32(in.Limit), int32(in.Offset))
	if err != nil {
		return nil, errors.Wrapf(err, "search msgs failed:%s", err.Error())
	}
	resp = new(pb.SearchHistoryMessageResp)
	for _, data := range datas {
		var msg pb.MessageWithSeq
		err = copier.Copy(&msg, &data)
		if err != nil {
			return nil, errors.Wrapf(err, "[search messages err]:copy form data to msg failed:%s", err.Error())
		}
		resp.Msg = append(resp.Msg, &msg)
	}
	return
}
