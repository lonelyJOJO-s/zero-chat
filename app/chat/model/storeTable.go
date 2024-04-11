package model

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

var _ StoreTableModel = (*customStoreTableModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	StoreTableModel interface {
		storeTableModel
	}

	customStoreTableModel struct {
		*defaultStoreModel
	}
)

type (
	storeTableModel interface {
		PutRow(primaryKeys map[string]any, columns map[string]any) error
	}

	defaultStoreModel struct {
		client *tablestore.TableStoreClient
		table  string
	}
)

var _ storeTableModel = (*defaultStoreModel)(nil)

// NewUsersModel returns a model for the database table.
func NewStoreTableModel(table string, client *tablestore.TableStoreClient) StoreTableModel {
	return &customStoreTableModel{
		defaultStoreModel: newStoreTableModel(table, client),
	}
}

func newStoreTableModel(table string, client *tablestore.TableStoreClient) *defaultStoreModel {
	return &defaultStoreModel{
		table:  table,
		client: client,
	}
}

func (m *defaultStoreModel) PutRow(primaryKeys map[string]any, columns map[string]any) error {
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
