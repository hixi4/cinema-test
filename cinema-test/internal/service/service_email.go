package service

import (
	"fmt"
	"net/smtp"
)

// EmailService відповідає за надсилання електронних листів через SMTP-сервер
type EmailService struct {
	smtpHost string
	smtpPort string
	username string
	password string
}

// NewEmailService створює новий екземпляр EmailService
func NewEmailService(smtpHost, smtpPort, username, password string) *EmailService {
	return &EmailService{
		smtpHost: smtpHost,
		smtpPort: smtpPort,
		username: username,
		password: password,
	}
}

// SendEmail надсилає електронний лист
func (e *EmailService) SendEmail(to, subject, body string) error {
	from := e.username
	pass := e.password

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail(e.smtpHost+":"+e.smtpPort,
		smtp.PlainAuth("", from, pass, e.smtpHost),
		from, []string{to}, []byte(msg))

	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
