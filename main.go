package main

import (
	// kafka "go_alert/go_kafka"
	// "go_alert/util"
	// "log"
	"fmt"
	trend "go_alert/chart_trend"
)

var CryptoAPISrc = trend.APISource{
	Url:    "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest",
	Method: "GET",
}

func main() {
	data := trend.GetPrice("btc", CryptoAPISrc)
	fmt.Println(data)

	// cfg, err_cfg := util.LoadConfig(".")
	// <<<<<<< Send Email >>>>>>>
	// if error != nil{
	// 	log.Println("Error:", error)
	// }
	// emailSender := go_mail.NewGmailSender(cfg.EmailSenderAddress, cfg.EmailSenderPassword)
	// From := "vietpride295@gmail.com"
	// To := []string{"viettran295@gmail.com"}
	// Subject, Text := "Go test", "Test Go"
	// emailSender.SendEmail(From, To, Subject, []byte(Text))

	// <<<<<<< ETH >>>>>>
	// client, err := ethclient.DialContext(context.Background(), infuraURL)
	// if err != nil{
	// 	log.Println(err)
	// }
	// defer client.Close()
	// block, err := client.BlockByNumber(context.Background(), nil)
	// log.Print(block.Number())

	// <<<<<< Kafka >>>>>>
	// if err_cfg != nil {
	// 	log.Println("Can not load config info")
	// }
	// consumer := kafka.NewConsumer(cfg.KafkaBootstrapServer, cfg.ConsumerGroupid)
	// kafka.Kafka_consume(consumer, []string{"saigon"})
}
