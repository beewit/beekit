package mail

import "testing"
import "log"

func TestSender(t *testing.T) {
	sender := &MailSender{
		Username: "xxxxx@163.com",
		Password: "xx",
		Host:     "smtp.163.com",
		Port:     456}
	/*
		sender := &MailSender{
			Username: "xxxxx@qq.com",
			Password: "xx",
			Host:     "smtp.qq.com",
			Port:     465}
	*/
	if err := sender.SendMail("GOMAIL Test", "hello world!", []string{"ec.huyinghuan@gmail.com"}); err != nil {
		log.Fatal(err)
		t.Fail()
	} else {
		log.Println("发送成功！")
	}
}
