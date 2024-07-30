// static file upload to oss bucket
package oss

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var bucket *oss.Bucket

func init() {
	godotenv.Load()
	/// 从环境变量中获取访问凭证。运行本代码示例之前，请确保已设置环境变量OSS_ACCESS_KEY_ID和OSS_ACCESS_KEY_SECRET。
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New("https://oss-cn-shanghai.aliyuncs.com", "", "", oss.SetCredentialsProvider(&provider))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err = client.Bucket("chat-bkt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func Upload2Oss(file []byte, filename string) (err error) {
	err = bucket.PutObject(filename, bytes.NewReader(file))
	if err != nil {
		logrus.Info(err)
	}
	return
}
