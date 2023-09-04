package mail_alert

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

const(
	smtpAuthAddress = "smtp.gmail.com"
	smtpServer = "smtp.gmail.com:587"

)

type EmailSender interface{
	SendEmail(
		From string,
		To []string,
		Subject string, 
		Text []byte,
	) error
}

type GmailSender struct{
	FromEmailAddress string
	FromEmailPassword string
}

func NewGmailSender(address string, password string) *GmailSender{
	return &GmailSender{
		FromEmailAddress: address,
		FromEmailPassword: password,
	}
}

func (sender *GmailSender) SendEmail(
	from string, 
	to []string, 
	subject string, 
	text []byte,
) error{
	e := email.NewEmail()
	e.From = from
	e.To = to
	e.Subject = subject
	e.Text = text
	smtpAuth := smtp.PlainAuth("", sender.FromEmailAddress, sender.FromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServer, smtpAuth)
}