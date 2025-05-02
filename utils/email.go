package utils

import (
	"log/slog"

	"github.com/go-gomail/gomail"
)

type SmtpServer struct {
	Host     string
	Port     int
	Username string
	Password string
}

func (s SmtpServer) SetDialer() *gomail.Dialer {
	dialer := gomail.NewDialer(s.Host, s.Port, s.Username, s.Password)
	return dialer
}

func (s SmtpServer) SendMail(to []string, from string, subject string, msg []byte) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", string(msg))

	dialer := s.SetDialer()

	if err := dialer.DialAndSend(m); err != nil {
		slog.Info("Failed to send email", "error", err)
		return err
	}
	return nil
}
