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
	groupsFieldNames          = builder.RawFieldNames(&Groups{})
	groupsRows                = strings.Join(groupsFieldNames, ",")
	groupsRowsExpectAutoSet   = strings.Join(stringx.Remove(groupsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	groupsRowsWithPlaceHolder = strings.Join(stringx.Remove(groupsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUsercenterGroupsIdPrefix = "cache:usercenter:groups:id:"
)

type (
	groupsModel interface {
		Insert(ctx context.Context, data *Groups) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Groups, error)
		Update(ctx context.Context, data *Groups) error
		Delete(ctx context.Context, id int64) error
	}

	defaultGroupsModel struct {
		sqlc.CachedConn
		table string
	}

	Groups struct {
		Id        int64          `db:"id"`
		CreatedAt time.Time      `db:"created_at"`
		UpdatedAt time.Time      `db:"updated_at"`
		DeletedAt sql.NullTime   `db:"deleted_at"`
		Name      sql.NullString `db:"name"`
		Desc      sql.NullString `db:"desc"`
		Avatar    sql.NullString `db:"avatar"`
		OwnerId   int64          `db:"owner_id"`
	}
)

func newGroupsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultGroupsModel {
	return &defaultGroupsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`groups`",
	}
}

func (m *defaultGroupsModel) Delete(ctx context.Context, id int64) error {
	usercenterGroupsIdKey := fmt.Sprintf("%s%v", cacheUsercenterGroupsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, usercenterGroupsIdKey)
	return err
}

func (m *defaultGroupsModel) FindOne(ctx context.Context, id int64) (*Groups, error) {
	usercenterGroupsIdKey := fmt.Sprintf("%s%v", cacheUsercenterGroupsIdPrefix, id)
	var resp Groups
	err := m.QueryRowCtx(ctx, &resp, usercenterGroupsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupsRows, m.table)
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

func (m *defaultGroupsModel) Insert(ctx context.Context, data *Groups) (sql.Result, error) {
	usercenterGroupsIdKey := fmt.Sprintf("%s%v", cacheUsercenterGroupsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, groupsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.Name, data.Desc, data.Avatar, data.OwnerId)
	}, usercenterGroupsIdKey)
	return ret, err
}

func (m *defaultGroupsModel) Update(ctx context.Context, data *Groups) error {
	usercenterGroupsIdKey := fmt.Sprintf("%s%v", cacheUsercenterGroupsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, groupsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.Name, data.Desc, data.Avatar, data.OwnerId, data.Id)
	}, usercenterGroupsIdKey)
	return err
}

func (m *defaultGroupsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUsercenterGroupsIdPrefix, primary)
}

func (m *defaultGroupsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultGroupsModel) tableName() string {
	return m.table
}
