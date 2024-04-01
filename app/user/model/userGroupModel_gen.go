// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userGroupFieldNames          = builder.RawFieldNames(&UserGroup{})
	userGroupRows                = strings.Join(userGroupFieldNames, ",")
	userGroupRowsExpectAutoSet   = strings.Join(stringx.Remove(userGroupFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userGroupRowsWithPlaceHolder = strings.Join(stringx.Remove(userGroupFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUsercenterUserGroupIdPrefix            = "cache:usercenter:userGroup:id:"
	cacheUsercenterUserGroupUserIdGroupIdPrefix = "cache:usercenter:userGroup:userId:groupId:"
)

type (
	userGroupModel interface {
		Insert(ctx context.Context, data *UserGroup) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserGroup, error)
		FindOneByUserIdGroupId(ctx context.Context, userId int64, groupId int64) (*UserGroup, error)
		Update(ctx context.Context, data *UserGroup) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserGroupModel struct {
		sqlc.CachedConn
		table string
	}

	UserGroup struct {
		Id        int64        `db:"id"`
		GroupId   int64        `db:"group_id"`
		UserId    int64        `db:"user_id"`
		CreatedAt time.Time    `db:"created_at"`
		DeletedAt sql.NullTime `db:"deleted_at"`
	}
)

func newUserGroupModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserGroupModel {
	return &defaultUserGroupModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user_group`",
	}
}

func (m *defaultUserGroupModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	usercenterUserGroupIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserGroupIdPrefix, id)
	usercenterUserGroupUserIdGroupIdKey := fmt.Sprintf("%s%v:%v", cacheUsercenterUserGroupUserIdGroupIdPrefix, data.UserId, data.GroupId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, usercenterUserGroupIdKey, usercenterUserGroupUserIdGroupIdKey)
	return err
}

func (m *defaultUserGroupModel) FindOne(ctx context.Context, id int64) (*UserGroup, error) {
	usercenterUserGroupIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserGroupIdPrefix, id)
	var resp UserGroup
	err := m.QueryRowCtx(ctx, &resp, usercenterUserGroupIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userGroupRows, m.table)
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

func (m *defaultUserGroupModel) FindOneByUserIdGroupId(ctx context.Context, userId int64, groupId int64) (*UserGroup, error) {
	usercenterUserGroupUserIdGroupIdKey := fmt.Sprintf("%s%v:%v", cacheUsercenterUserGroupUserIdGroupIdPrefix, userId, groupId)
	var resp UserGroup
	err := m.QueryRowIndexCtx(ctx, &resp, usercenterUserGroupUserIdGroupIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `group_id` = ? limit 1", userGroupRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, groupId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserGroupModel) Insert(ctx context.Context, data *UserGroup) (sql.Result, error) {
	usercenterUserGroupIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserGroupIdPrefix, data.Id)
	usercenterUserGroupUserIdGroupIdKey := fmt.Sprintf("%s%v:%v", cacheUsercenterUserGroupUserIdGroupIdPrefix, data.UserId, data.GroupId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userGroupRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.GroupId, data.UserId, data.DeletedAt)
	}, usercenterUserGroupIdKey, usercenterUserGroupUserIdGroupIdKey)
	return ret, err
}

func (m *defaultUserGroupModel) Update(ctx context.Context, newData *UserGroup) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	usercenterUserGroupIdKey := fmt.Sprintf("%s%v", cacheUsercenterUserGroupIdPrefix, data.Id)
	usercenterUserGroupUserIdGroupIdKey := fmt.Sprintf("%s%v:%v", cacheUsercenterUserGroupUserIdGroupIdPrefix, data.UserId, data.GroupId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userGroupRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.GroupId, newData.UserId, newData.DeletedAt, newData.Id)
	}, usercenterUserGroupIdKey, usercenterUserGroupUserIdGroupIdKey)
	return err
}

func (m *defaultUserGroupModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUsercenterUserGroupIdPrefix, primary)
}

func (m *defaultUserGroupModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userGroupRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserGroupModel) tableName() string {
	return m.table
}
