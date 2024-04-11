package ws

import (
	"encoding/base64"
	"fmt"
	"log"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
)

func TestProtoGen(t *testing.T) {

	// 创建 Protobuf 消息实例并填充字段
	msg := Message{}
	msg.From = 5
	msg.Content = "Hello, this is a message."
	msg.SendTime = time.Now().UnixNano()
	msg.ContentType = 0 // 0 for text
	msg.To = 6
	msg.File = []byte{}
	msg.ChatType = 0 // o for signle 1 for group
	msg.Type = "heartbeat"

	// 序列化 Protobuf 消息为二进制数据
	fmt.Println(msg)
	data, err := proto.Marshal(&msg)
	if err != nil {
		log.Fatalf("Failed to marshal message: %v", err)
	}
	res := base64Encode(data)
	fmt.Println(res)

}

func base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
