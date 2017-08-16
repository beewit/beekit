package mail

import (
	"crypto/tls"
	"fmt"
	"net/mail"
	"net/smtp"
	"strconv"
	"strings"
)

//MailSender 邮件发送人
type MailSender struct {
	Username string
	Password string
	Host     string
	Port     int
}

//SendMail 发送邮件
//subject 邮件标题
//content 邮件内容
//contact 联系人
func (sender *MailSender) SendMail(subject string, content string, contact []string) error {
	if len(contact) == 0 {
		return fmt.Errorf("联系人不能为空. Contact cannot be empty!\n")
	}
	from := mail.Address{"", sender.Username}
	to := mail.Address{"", strings.Join(contact, ";")}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	// Build Email Content
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + content

	auth := smtp.PlainAuth("", sender.Username, sender.Password, sender.Host)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         sender.Host}

	conn, err := tls.Dial("tcp", sender.Host+":"+strconv.Itoa(sender.Port), tlsconfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, sender.Host)
	if err != nil {
		return err
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		return err
	}
	if err = c.Mail(from.Address); err != nil {
		return err
	}
	if err = c.Rcpt(to.Address); err != nil {
		return err
	}
	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	c.Quit()
	return nil
}
