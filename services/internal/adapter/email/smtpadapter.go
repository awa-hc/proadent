package email

import (
	"os"
	"services/internal/application/port"
	"strconv"

	"gopkg.in/gomail.v2"
)

type SMTPAdapter struct {
	Host     string
	Port     int
	Email    string
	Password string
}

func SetupSMTPAdapter() *SMTPAdapter {
	smtpHost := os.Getenv("EMAIL_HOST")
	smtpPort, _ := strconv.Atoi(os.Getenv("EMAIL_PORT")) // Asegúrate de que EMAIL_PORT exista y sea un número válido.
	smtpUsername := os.Getenv("EMAIL_FROM")
	smtpPassword := os.Getenv("EMAIL_PASSWORD")

	return NewSMTPAdapter(smtpHost, smtpPort, smtpUsername, smtpPassword)
}

func NewSMTPAdapter(host string, port int, email string, password string) *SMTPAdapter {
	return &SMTPAdapter{
		Host:     host,
		Port:     port,
		Email:    email,
		Password: password,
	}
}

func (adapter *SMTPAdapter) SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", adapter.Email)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(adapter.Host, adapter.Port, adapter.Email, adapter.Password)

	// Enviar el email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

var _ port.EmailSender = &SMTPAdapter{}
