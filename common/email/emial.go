package email

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
)

const (
	LoginType = iota
	RegisterType
)

var Msg map[int]string = map[int]string{
	LoginType:    "Login Certification",
	RegisterType: "Account Register",
}
var p *email.Pool

func init() {
	var err error
	p, err = email.NewPool(
		"smtp.qq.com:25",
		4,
		smtp.PlainAuth("", "joe8273@qq.com", "mzcvoelypvmbhbei", "smtp.qq.com"),
	)
	if err != nil {
		log.Fatal("failed to create pool:", err)
	}
}

func SendEmail(addrs []string, emailType int) (string, error) {
	e := email.NewEmail()
	e.From = "joe8273@qq.com"
	e.To = addrs
	code := GenerateCaptcha(6)
	switch emailType {
	case LoginType:
		e.Subject = Msg[LoginType]
		e.Text = []byte(fmt.Sprintf("Your login in vertify code is:%s \n", code))
	case RegisterType:
	default:
	}
	go p.Send(e, 10*time.Second)
	return code, nil
}

func GenerateCaptcha(length int) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
