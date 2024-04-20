package ctxdata

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContextKey string

const CtxKeyJwtUserId ContextKey = "jwtUserId"

// CtxKeyJwtUserId get uid from ctx

func GetUidFromCtx(ctx context.Context) int64 {
	var uid int64
	fmt.Println(ctx.Value(CtxKeyJwtUserId))
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
		}
	}
	return uid
}
