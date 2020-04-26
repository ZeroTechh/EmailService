package email

import (
	"net/smtp"

	"github.com/ZeroTechh/hades"
)

var (
	config    = hades.GetConfig("main.yaml", []string{"config", "../config", "../../config"})
	sendMails = config.Map("service").Bool("sendMails")

	smtpConfig = config.Map("smtp")
	host       = smtpConfig.Str("host")
	port       = smtpConfig.Str("post")
	email      = smtpConfig.Str("email")
	password   = smtpConfig.Str("password")
	hostname   = host + ":" + port
)

// New creates a new Email.
func New() *Email {
	e := Email{}
	e.init()
	return &e
}

// Email will send emails.
type Email struct {
	auth smtp.Auth
}

// init will initialize.
func (e *Email) init() {
	e.auth = smtp.PlainAuth("", email, password, host)
}

// Simple will send a simple text email.
func (e Email) Simple(text, to string) error {
	if !sendMails {
		return nil
	}

	return smtp.SendMail(
		hostname, e.auth, email, []string{to}, []byte(text))
}
