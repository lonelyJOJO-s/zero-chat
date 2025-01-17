// Code generated by goctl. DO NOT EDIT.
package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	groupsFieldNames          = builder.RawFieldNames(&Groups{})
	groupsRows                = strings.Join(groupsFieldNames, ",")
	groupsRowsExpectAutoSet   = strings.Join(stringx.Remove(groupsFieldNames, "`id`", "`deleted_at`"), ",")
	groupsRowsWithPlaceHolder = strings.Join(stringx.Remove(groupsFieldNames, "`id`", "`created_at`, `deleted_at`"), "=?,") + "=?"

	cacheUsercenterGroupsIdPrefix = "cache:usercenter:groups:id:"
)

type (
	groupsModel interface {
		Insert(ctx context.Context, data *Groups) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Groups, error)
		Update(ctx context.Context, data *Groups, session sqlx.Session) error
		Delete(ctx context.Context, id int64) error
		DeleteSoft(ctx context.Context, id int64) error
		// all id 
		FindAllByUserId(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]int64, error)
		// all struct
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Groups, error)
		SelectBuilder() squirrel.SelectBuilder
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
		LastMessageTime time.Time 	 `db:"last_message_time"`
	}
)

var _ groupsModel = (*defaultGroupsModel)(nil)

func newGroupsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultGroupsModel {
	return &defaultGroupsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`groups`",
	}
}

func (m *defaultGroupsModel) Delete(ctx context.Context, id int64) error {
	usercenterGroupsIdKey := fmt.Sprintf("%s%v", cacheUsercenterGroupsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ? and deleted_at is null", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, usercenterGroupsIdKey)
	return err
}

func (m *defaultGroupsModel) DeleteSoft(ctx context.Context, id int64) error {
	usercenterGroupsIdKey := fmt.Sprintf("%s%v", cacheUsercenterGroupsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set deleted_at= ?, owner_id= 0 where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, time.Now(), id)
	}, usercenterGroupsIdKey)
	fmt.Println(fmt.Sprintf("update %s set deleted_at= ?, owner_id= 0 where `id` = ?", m.table))
	return err
}

func (m *defaultGroupsModel) FindOne(ctx context.Context, id int64) (*Groups, error) {
	usercenterGroupsIdKey := fmt.Sprintf("%s%v", cacheUsercenterGroupsIdPrefix, id)
	var resp Groups
	err := m.QueryRowCtx(ctx, &resp, usercenterGroupsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and owner_id != 0 and deleted_at is null limit 1", groupsRows, m.table)
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

func (m *defaultGroupsModel) FindAllByUserId(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]int64, error) {
	rowBuilder = rowBuilder.Columns("`id`")

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where(squirrel.Eq{"deleted_at": nil}).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []int64
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultGroupsModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Groups, error) {
	rowBuilder = rowBuilder.Columns(groupsRows)

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.Where(squirrel.Eq{"deleted_at": nil}).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Groups
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultGroupsModel) Insert(ctx context.Context, data *Groups) (sql.Result, error) {
	usercenterGroupsIdKey := fmt.Sprintf("%s%v", cacheUsercenterGroupsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, groupsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.Name, data.Desc, data.Avatar, data.OwnerId)
	}, usercenterGroupsIdKey)
	return ret, err
}

func (m *defaultGroupsModel) Update(ctx context.Context, data *Groups, session sqlx.Session) error {
	usercenterGroupsIdKey := fmt.Sprintf("%s%v", cacheUsercenterGroupsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set updated_at =?, `name` =?, `desc` =?, avatar=?, owner_id=? where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, data.UpdatedAt, data.Name.String, data.Desc.String, data.Avatar.String, data.OwnerId, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.UpdatedAt, data.Name.String, data.Desc.String, data.Avatar.String, data.OwnerId, data.Id)
	}, usercenterGroupsIdKey)
	return err
}

func (m *defaultGroupsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUsercenterGroupsIdPrefix, primary)
}

func (m *defaultGroupsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and owner_id != 0 limit 1", groupsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultGroupsModel) tableName() string {
	return m.table
}


func (m *defaultGroupsModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

