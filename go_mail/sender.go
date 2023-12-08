package go_mail

import (
	"go_alert/util"
	"log"
	"net/smtp"
	"strconv"
)

const (
	smtpAuthAddress = "smtp.gmail.com"
	smtpServer      = "smtp.gmail.com:587"
)

func CreateAlertMsg(symbol string, typ string, value float64) (string, string) {
	Subject := symbol + " ALERT !" 
	Msg := typ + " of " + symbol + " is " + strconv.FormatFloat(float64(value), 'f', -1, 64)
	return Subject, Msg
}

func SendEmail(
	from string,
	to []string,
	subject string,
	text string,
) {
	cfg, err_cfg := util.LoadConfig(".")
	if err_cfg != nil {
		log.Panicln("Error while loading config")
	}

	msg := "Subject: " + subject + "\r\n" + text
	err := smtp.SendMail(smtpServer,
		smtp.PlainAuth("", cfg.EmailSenderAddress, cfg.EmailSenderPassword, smtpAuthAddress),
		from, to, []byte(msg))
	if err != nil {
		log.Panicln("Error while sending email")
	}
}
