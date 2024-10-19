package main

import (
	"fmt"
	"net/http"
)

func HttpPost(url string, client *http.Client) error {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjg5NjIzNTQsImlhdCI6MTcyODc4OTU1NCwiand0VXNlcklkIjo2fQ.GKjJ-cO2_Oo7VChkAo3Bf2OtdVvoZpuv5RHjz3PdpoI"

	// 创建请求
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	// 设置 Authorization 头
	req.Header.Set("Authorization", "Bearer "+token)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create user, status code: %d", resp.StatusCode)
	}

	return nil
}

// func main() {

// 	client := &http.Client{}
// 	var wg sync.WaitGroup
// 	for k := 656; k < 10006; k++ {
// 		url := fmt.Sprintf("http://localhost:8088/usercenter/api/v1/user/friend/%d", k)
// 		go func() {
// 			wg.Add(1)
// 			HttpPost(url, client)
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()

// }
