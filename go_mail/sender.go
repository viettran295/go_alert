package go_mail

import (
	"go_alert/util"
	"log"
	"net/smtp"
)

const (
	smtpAuthAddress = "smtp.gmail.com"
	smtpServer      = "smtp.gmail.com:587"
)

func SendEmail(
	from string,
	to []string,
	subject string,
	text string,
) {
	cfg, err_cfg := util.LoadConfig(".")
	if err_cfg != nil {
		log.Println("Error while loading config")
	}

	msg := "Subject: " + subject + "\r\n" + text
	err := smtp.SendMail(smtpServer,
		smtp.PlainAuth("", cfg.EmailSenderAddress, cfg.EmailSenderPassword, smtpAuthAddress),
		from, to, []byte(msg))
	if err != nil {
		log.Println("Error while sending email")
	}
}
