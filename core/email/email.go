package email

import (
	"net/smtp"

	"github.com/ZeroTechh/hades"
)

var (
	config     = hades.GetConfig("main.yaml", []string{"config", "../config", "../../config"})
	smtpConfig = config.Map("smtp")
	sendMails  = config.Map("service").Bool("sendMails")
)

// Email will handle sending of emails
type Email struct {
	auth smtp.Auth
}

// Init initializes
func (email *Email) Init() {
	email.auth = smtp.PlainAuth(
		"",
		smtpConfig.Str("email"),
		smtpConfig.Str("password"),
		smtpConfig.Str("host"),
	)
}

// SendSimpleEmail Is Used To Send A Basic Text Email
func (email Email) SendSimpleEmail(text, toEmail string) error {
	if !sendMails {
		// exit the function if sendMails is false. usually used in debug mode
		return nil
	}

	return smtp.SendMail(
		smtpConfig.Str("host")+":"+smtpConfig.Str("port"),
		email.auth,
		smtpConfig.Str("email"),
		[]string{toEmail},
		[]byte(text),
	)
}
