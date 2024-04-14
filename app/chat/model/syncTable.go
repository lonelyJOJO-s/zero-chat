package model

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

var _ SyncTableModel = (*customSyncTableModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	SyncTableModel interface {
		syncTableModel
	}

	customSyncTableModel struct {
		*defaultSyncModel
	}
)

type (
	syncTableModel interface {
		PutRow(primaryKeys map[string]any, columns map[string]any) error
	}

	defaultSyncModel struct {
		client *tablestore.TableStoreClient
		table  string
	}
)

var _ syncTableModel = (*defaultSyncModel)(nil)

func NewSyncTableModel(table string, client *tablestore.TableStoreClient) SyncTableModel {
	return &customSyncTableModel{
		defaultSyncModel: newSyncTableModel(table, client),
	}
}

func newSyncTableModel(table string, client *tablestore.TableStoreClient) *defaultSyncModel {
	return &defaultSyncModel{
		table:  table,
		client: client,
	}
}

func (m *defaultSyncModel) PutRow(primaryKeys map[string]any, columns map[string]any) error {
	columns["file"] = []byte("test use")
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = m.table
	putPk := new(tablestore.PrimaryKey)
	for primaryKey, primaryVal := range primaryKeys {
		putPk.AddPrimaryKeyColumn(primaryKey, primaryVal)
	}
	putRowChange.PrimaryKey = putPk
	for column, val := range columns {
		putRowChange.AddColumn(column, val)
	}
	// putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = putRowChange
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	_, err := m.client.PutRow(putRowRequest)
	return err
}
