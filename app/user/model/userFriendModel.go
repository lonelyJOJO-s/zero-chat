package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserFriendModel = (*customUserFriendModel)(nil)

type (
	// UserFriendModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserFriendModel.
	UserFriendModel interface {
		userFriendModel
	}

	customUserFriendModel struct {
		*defaultUserFriendModel
	}
)

// NewUserFriendModel returns a model for the database table.
func NewUserFriendModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserFriendModel {
	return &customUserFriendModel{
		defaultUserFriendModel: newUserFriendModel(conn, c, opts...),
	}
}
