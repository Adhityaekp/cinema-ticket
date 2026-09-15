package service

import (
	"fmt"
	"net/smtp"
)

type EmailService struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
	BaseURL  string
}

func NewEmailService(
	host string,
	port string,
	username string,
	password string,
	from string,
	baseURL string,
) *EmailService {

	return &EmailService{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
		BaseURL:  baseURL,
	}
}

func (s *EmailService) SendVerificationEmail(
	to string,
	token string,
) error {

	verifyURL := fmt.Sprintf(
		"%s/api/auth/verify-email?token=%s",
		s.BaseURL,
		token,
	)

	subject := "Cinema Ticket - Verifikasi Email"

	body := fmt.Sprintf(`
Halo,

Terima kasih telah melakukan registrasi.

Silakan klik link berikut untuk memverifikasi email Anda:

%s

Link ini digunakan untuk mengaktifkan akun Cinema Ticket.

Terima kasih.
`, verifyURL)

	message := []byte(
		"From: " + s.From + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			body,
	)

	auth := smtp.PlainAuth(
		"",
		s.Username,
		s.Password,
		s.Host,
	)

	return smtp.SendMail(
		s.Host+":"+s.Port,
		auth,
		s.From,
		[]string{to},
		message,
	)
}
