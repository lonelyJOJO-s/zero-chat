// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
	"zero-chat/common/xerr"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFriendFieldNames          = builder.RawFieldNames(&UserFriend{})
	userFriendRows                = strings.Join(userFriendFieldNames, ",")
	userFriendRowsExpectAutoSet   = strings.Join(stringx.Remove(userFriendFieldNames, "`id`", "`create_time`", "`deleted_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userFriendRowsWithPlaceHolder = strings.Join(stringx.Remove(userFriendFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUsercenterUserFriendIdPrefix             = "cache:usercenter:userFriend:id:"
	cacheUsercenterUserFriendUserIdFriendIdPrefix = "cache:usercenter:userFriend:userId:friendId:"
)

type (
	userFriendModel interface {
		Insert(ctx context.Context, data *UserFriend) (error)
		FindOne(ctx context.Context, id int64) (*UserFriend, error)
		FindOneByUserIdFriendId(ctx context.Context, userId int64, friendId int64) (*UserFriend, error)
		Update(ctx context.Context, data *UserFriend) error
		Delete(ctx context.Context, id int64) error
		GetFriendIds(ctx context.Context, id int64) ([]int, error)
		DeleteSoft(ctx context.Context, uf UserFriend) error
		DeleteBuilder() squirrel.DeleteBuilder
		UpdateBuilder() squirrel.UpdateBuilder
	}

	defaultUserFriendModel struct {
		sqlc.CachedConn
		table string
	}

	UserFriend struct {
		Id        int64        `db:"id"`
		UserId    int64        `db:"user_id"`
		FriendId  int64        `db:"friend_id"`
		Uuid      string       `db:"uuid"`
		CreatedAt time.Time    `db:"created_at"`
		DeletedAt sql.NullTime `db:"deleted_at"`
	}
)

var _ userFriendModel = (*defaultUserFriendModel)(nil)

func newUserFriendModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserFriendModel {
	return &defaultUserFriendModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user_friend`",
	}
}

func (m *defaultUserFriendModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	usercenterUserFriendIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserFriendIdPrefix, id)
	usercenterUserFriendUserIdFriendIdKey := fmt.Sprintf("%s%v:%v", cacheUsercenterUserFriendUserIdFriendIdPrefix, data.UserId, data.FriendId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, usercenterUserFriendIdKey, usercenterUserFriendUserIdFriendIdKey)
	return err
}

func (m *defaultUserFriendModel) DeleteSoft(ctx context.Context, uf UserFriend) error {
	now := time.Now()
	builder := m.UpdateBuilder().Where(squirrel.Or{
		squirrel.Eq{"user_id": uf.UserId, "friend_id": uf.FriendId},
		squirrel.Eq{"user_id": uf.FriendId, "friend_id": uf.UserId},
	}).Set("deleted_at", now)
	query, values, err := builder.ToSql()
	if err != nil {
		return err
	}
	_, err = m.ExecNoCacheCtx(ctx, query, values...)
	if err != nil {
		return err
	}
	return nil
}


func (m *defaultUserFriendModel) GetFriendIds(ctx context.Context, id int64) (ids []int, err error) {
	query := fmt.Sprintf("select friend_id from %s where `user_id` = ? and deleted_at is null", m.table)
	err = m.QueryRowsNoCacheCtx(ctx, &ids, query, id)
	if err != nil {
		return nil, err
	}
	return 
}

func (m *defaultUserFriendModel) FindOne(ctx context.Context, id int64) (*UserFriend, error) {
	usercenterUserFriendIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserFriendIdPrefix, id)
	var resp UserFriend
	err := m.QueryRowCtx(ctx, &resp, usercenterUserFriendIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and deleted_at is null limit 1", userFriendRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserFriendModel) FindOneByUserIdFriendId(ctx context.Context, userId int64, friendId int64) (*UserFriend, error) {
	var resp UserFriend
	
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `friend_id` = ? and deleted_at is null limit 1", userFriendRows, m.table)
	if err := m.QueryRowNoCacheCtx(ctx, &resp, query, userId, friendId); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserFriendModel) Insert(ctx context.Context, data *UserFriend) (error) {
	
	uf, err := m.FindOneByUserIdFriendId(ctx, data.UserId, data.FriendId)
	if err != nil && err != ErrNotFound{
		return err
	}
	if uf != nil {
		return xerr.NewErrCodeMsg(xerr.INSERT_ALREADY_EXSIT, xerr.MapErrMsg(xerr.INSERT_ALREADY_EXSIT))
	}
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, userFriendRowsExpectAutoSet)
	now := time.Now()
	uuid := uuid.NewString()
	_, err = m.ExecNoCacheCtx(ctx, query, data.UserId, data.FriendId, uuid, now)
	if err != nil {
		return  err
	}
	_, err = m.ExecNoCacheCtx(ctx, query, data.FriendId, data.UserId , uuid, now)
	if err != nil {
		return  err
	}
	return  nil
}

func (m *defaultUserFriendModel) DeleteAll(ctx context.Context, builder squirrel.DeleteBuilder) (error) {
	return nil
}

func (m *defaultUserFriendModel) Update(ctx context.Context, newData *UserFriend) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	usercenterUserFriendIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserFriendIdPrefix, data.Id)
	usercenterUserFriendUserIdFriendIdKey := fmt.Sprintf("%s%v:%v", cacheUsercenterUserFriendUserIdFriendIdPrefix, data.UserId, data.FriendId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and deleted_at is null", m.table, userFriendRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserId, newData.FriendId, newData.DeletedAt, newData.Id)
	}, usercenterUserFriendIdKey, usercenterUserFriendUserIdFriendIdKey)
	return err
}

func (m *defaultUserFriendModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUsercenterUserFriendIdPrefix, primary)
}

func (m *defaultUserFriendModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and deleted_at is null limit 1", userFriendRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserFriendModel) tableName() string {
	return m.table
}

func (m *defaultUserFriendModel) DeleteBuilder() squirrel.DeleteBuilder {
	return squirrel.Delete(m.table)
}

func (m *defaultUserFriendModel) UpdateBuilder() squirrel.UpdateBuilder {
	return squirrel.Update(m.table)
}
