package svc

import (
	"time"
	"zero-chat/app/user/cmd/rpc/internal/config"
	"zero-chat/app/user/model"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	GroupModel  model.GroupsModel
	UserModel   model.UsersModel
	UserFriend  model.UserFriendModel
	UserGroup   model.UserGroupModel
	SqlConn     sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		SqlConn: sqlConn,
		Config:  c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		UserModel:  model.NewUsersModel(sqlConn, c.Cache, cache.WithExpiry(time.Hour*24*3)), // default is seven days
		GroupModel: model.NewGroupsModel(sqlConn, c.Cache, cache.WithExpiry(time.Hour*24*3)),
		UserFriend: model.NewUserFriendModel(sqlConn, c.Cache, cache.WithExpiry(time.Hour*24*3)),
		UserGroup:  model.NewUserGroupModel(sqlConn, c.Cache, cache.WithExpiry(time.Hour*24*3)),
	}
}
