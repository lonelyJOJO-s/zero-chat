package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore/search"
)

func QueryWithPrimaryKeyAndContent(client *tablestore.TableStoreClient, tableName string, indexName string, keyVal string, limit int32) {
	// 构建 BoolQuery
	boolQuery := &search.BoolQuery{}

	// 主键查询
	primaryKeyQuery := &search.TermQuery{
		FieldName: "TimelineId",
		Term:      keyVal, // 你的主键值
	}
	boolQuery.MustQueries = append(boolQuery.MustQueries, primaryKeyQuery)

	// 模糊匹配 Content 字段，类似 SQL 中的 LIKE '%yyy%'
	contentWildcardQuery := &search.WildcardQuery{
		FieldName: "Content",
		Value:     "*nih*", // 模糊匹配值
	}
	boolQuery.MustQueries = append(boolQuery.MustQueries, contentWildcardQuery)

	// 创建搜索请求
	searchRequest := &tablestore.SearchRequest{}
	searchRequest.SetTableName(tableName)
	searchRequest.SetIndexName(indexName)

	searchQuery := search.NewSearchQuery()
	searchQuery.SetQuery(boolQuery)
	searchQuery.SetLimit(limit)
	searchRequest.SetSearchQuery(searchQuery)
	searchRequest.SetColumnsToGet(&tablestore.ColumnsToGet{
		ReturnAll: true,
	})

	// 执行查询
	searchResponse, err := client.Search(searchRequest)
	if err != nil {
		fmt.Printf("Search error: %v", err)
		return
	}

	// 打印搜索结果
	fmt.Println("IsAllSuccess: ", searchResponse.IsAllSuccess)
	fmt.Println("TotalCount: ", searchResponse.TotalCount)
	fmt.Println("RowsSize: ", len(searchResponse.Rows))
	for _, row := range searchResponse.Rows {
		jsonData, err := json.Marshal(row.Columns)
		if err != nil {
			panic(err)
		}
		var data CustomData
		err = json.Unmarshal([]byte(jsonData), &data)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Parsed Data: %+v\n", data)
	}

}

type Column struct {
	ColumnName string      `json:"ColumnName"`
	Value      interface{} `json:"Value"`
}

type CustomData struct {
	File        string `column:"Attr_File"`
	From        int64  `column:"Attr_Sender"`
	SendTime    int64  `column:"Timestamp"`
	Content     string `column:"Content"`
	ContentType int64  `column:"Attr_MsgType"`
	ChatType    int64  `column:"Attr_Type"`
}

// 自定义的 UnmarshalJSON 方法
func (cd *CustomData) UnmarshalJSON(data []byte) error {
	// 解析为 Column 的 slice
	var columns []Column
	if err := json.Unmarshal(data, &columns); err != nil {
		return err
	}

	// 通过反射获取 CustomData 的字段信息
	v := reflect.ValueOf(cd).Elem()
	t := v.Type()

	// 遍历所有 Column 并匹配对应的 struct 标签
	for _, column := range columns {
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get("column")
			if tag == column.ColumnName {
				// 根据字段类型来设置值
				switch v.Field(i).Kind() {
				case reflect.String:
					if val, ok := column.Value.(string); ok {
						v.Field(i).SetString(val)
					}
				case reflect.Int32, reflect.Int64:
					// 处理数字类型（JSON 中数字默认是 float64）
					if val, ok := column.Value.(float64); ok {
						v.Field(i).SetInt(int64(val))
					}
				}
			}
		}
	}
	return nil
}

func main() {

}
