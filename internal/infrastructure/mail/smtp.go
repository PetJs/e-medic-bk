// Package mail provides email sending implementations.
package mail

import (
	"context"

	"emedic-bk/internal/application/port"
)

// SMTPMailer implements port.Mailer using SMTP.
type SMTPMailer struct {
	host     string
	port     int
	username string
	password string
	from     string
}

// NewSMTPMailer creates a new SMTP mailer.
func NewSMTPMailer(host string, port int, username, password, from string) port.Mailer {
	return &SMTPMailer{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     from,
	}
}

func (m *SMTPMailer) SendPasswordReset(ctx context.Context, email, token string) error {
	// TODO: Implement password reset email
	return nil
}

func (m *SMTPMailer) SendWelcome(ctx context.Context, email, name string) error {
	// TODO: Implement welcome email
	return nil
}

func (m *SMTPMailer) SendSubscriptionConfirmation(ctx context.Context, email, planName string) error {
	// TODO: Implement subscription confirmation email
	return nil
}
