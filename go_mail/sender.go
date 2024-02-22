package go_mail

import (
	"go_alert/util"
	"log"
	"net/smtp"
	"strconv"
	"time"
)

const (
	smtpAuthAddress = "smtp.gmail.com"
	smtpServer      = "smtp.gmail.com:587"
)

func CreateAlertMsg(symbol string, typ string, value float64) {
	t := time.Now()
	Subject := symbol + " ALERT !"
	Msg := typ + " of " + symbol + " is " + strconv.FormatFloat(float64(value), 'f', -1, 64) + "\n" + t.Format(time.RFC3339)
	SendEmail(Subject, Msg)
}

func SendEmail(
	subject string,
	text string,
) {
	var From string = "vietpride295@gmail.com"
	var To = []string{"viettran295@gmail.com"}
	cfg, err_cfg := util.LoadConfig(".")
	if err_cfg != nil {
		log.Panicln("Error while loading config")
	}

	msg := "Subject: " + subject + "\r\n" + text
	err := smtp.SendMail(smtpServer,
		smtp.PlainAuth("", cfg.EmailSenderAddress, cfg.EmailSenderPassword, smtpAuthAddress),
		From, To, []byte(msg))
	if err != nil {
		log.Panicln("Error while sending email")
	}
}
