package model

import (
	"testing"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func TestTable(t *testing.T) {
	putRowRequest := new(tablestore.PutRowRequest)
	putRowChange := new(tablestore.PutRowChange)
	putRowChange.TableName = "im_xxx"
	putPk := new(tablestore.PrimaryKey)
	putPk.AddPrimaryKeyColumn("timeline_id", "pk1value1")
	putPk.AddPrimaryKeyColumn("sequence_id", int64(215646))
	putRowChange.PrimaryKey = putPk
	putRowChange.AddColumn("content", "anim")
	putRowChange.AddColumn("msg_type", int64(0))
	putRowChange.AddColumn("send_time", int64(111))
	putRowChange.AddColumn("sender", int64(111))
	putRowChange.AddColumn("type", int64(111))
	putRowChange.AddColumn("file", []byte("test"))
	putRowChange.AddColumn("conversation", "135465")
	putRowChange.SetCondition(tablestore.RowExistenceExpectation_IGNORE)
	putRowRequest.PutRowChange = putRowChange
	// _, err := client.PutRow(putRowRequest)

	// if err != nil {
	// 	fmt.Println("putrow failed with error:", err)
	// } else {
	// 	fmt.Println("putrow finished")
	// }
}
