package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/exp/rand"
)

type User struct {
	UserName        string `json:"username"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	ConfirmPassword string `json:"confirm_password"`
	Password        string `json:"password"`
	Sex             int8   `json:"sex"`
}

func createUser(client *http.Client, url string, user User) error {
	// 将用户结构体转换为 JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// 创建 POST 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(userJSON))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create user, status code: %d", resp.StatusCode)
	}

	return nil
}

func generatePhoneNumber() string {

	// 手机号码前缀以 1 开头，第二位通常是 3-9
	prefixes := []int{3, 5, 6, 7, 8, 9}
	secondDigit := prefixes[rand.Intn(len(prefixes))]

	// 生成剩下的 9 位随机数字
	phoneNumber := fmt.Sprintf("1%d", secondDigit)
	for i := 0; i < 9; i++ {
		phoneNumber += fmt.Sprintf("%d", rand.Intn(10))
	}

	return phoneNumber
}

// func main() {
// 	url := "http://localhost:8088/usercenter/api/v1/user/register"
// 	client := &http.Client{}
// 	var wg sync.WaitGroup
// 	for k := 0; k < 100; k++ {
// 		wg.Add(1)
// 		go func(k int) {
// 			for i := 1; i <= 100; i++ {
// 				user := User{
// 					UserName:        "test_user" + strconv.Itoa(100*k+i),
// 					Email:           "user" + strconv.Itoa(100*k+i) + "@test.com",
// 					Password:        "123456",
// 					ConfirmPassword: "123456",
// 					Sex:             0,
// 					Phone:           generatePhoneNumber(),
// 				}

// 				err := createUser(client, url, user)
// 				if err != nil {
// 					fmt.Println("Error creating user:", err)
// 					continue
// 				}

// 				fmt.Println("Successfully created user", i)
// 			}
// 			wg.Done()
// 		}(k)
// 	}
// 	wg.Wait()
// }
