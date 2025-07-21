package email

import (
	"Ai-Novel/common/zlog"
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Sender struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewEmailSender(host string, port int, username, password string) Sender {
	return Sender{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

func (s *Sender) Send(to []string, subject string, message string) error {
	// 1. 连接SMTP服务器
	host := s.Host
	port := s.Port
	userName := s.Username
	password := s.Password

	// 2. 构建邮件对象
	m := gomail.NewMessage()
	m.SetHeader("From", userName)   // 发件人
	m.SetHeader("To", to...)        // 收件人
	m.SetHeader("Subject", subject) // 主题
	m.SetBody("text/html", message) // 正文

	d := gomail.NewDialer(
		host,
		port,
		userName,
		password,
	)
	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		zlog.Errorf("邮件发送失败：%v", err)
		return err
	}
	return nil
}
