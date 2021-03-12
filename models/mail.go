package models

import (
	"EasyTutor/utils/logger"
	"github.com/beego/beego/v2/server/web"
	"net/smtp"
)

type Mail struct {
	To 		string `json:"to"`
	Subject string `json:"subject"`
	Msg     string `json:"msg"`
}

var (
	EmailFrom string
	EmailAuth smtp.Auth
)

func (m *Mail) Send(receiver string) error {
	to := "To: " + m.To +"\r\n"
	subject := "Subject: " + m.Subject  +"\r\n"
	message := []byte(to + subject + "\n" + m.Msg + "\r\n")
	hostPort, err := web.AppConfig.String("gmail::host_port")
	if err != nil {
		return err
	}
	err = smtp.SendMail(hostPort, EmailAuth, EmailFrom, []string{receiver}, message)
	if err != nil {
		logger.Error("[Error sending email] user receive = %v error = %v", receiver, err)
	}
	return err
}

func init() {
	emailHost := "smtp.gmail.com"
	EmailFrom = "ltdai2468@gmail.com"
	emailPassword := "dajape0#"
	EmailAuth = smtp.PlainAuth("", EmailFrom, emailPassword, emailHost)
}