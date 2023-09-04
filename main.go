package main

import (
	"log"
	"main/mail_alert"
	"main/util"
)

func main() {
	cfg, error := util.LoadConfig(".")
	if error != nil{
		log.Println("Error:", error)
	}
	emailSender := mail_alert.NewGmailSender(cfg.EmailSenderAddress, cfg.EmailSenderPassword)
	From := "vietpride295@gmail.com"
	To := []string{"viettran295@gmail.com"}
	Subject, Text := "Go test", "Test Go"
	emailSender.SendEmail(From, To, Subject, []byte(Text))

}
