package main

import (
	// kafka "go_alert/go_kafka"
	"go_alert/go_mail"
	"go_alert/processor"
	"go_alert/req"
	"log"
	"math"
	"time"
)

var CryptoAPISrc = req.APISource{
	Url:    "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest",
	Method: "GET",
}

func main() {
	start := time.Now()
	ch := make(chan req.CryptoAPIResponse)
	CryptoSym := []string{"BTC", "ETH", "SOL", "XRP", "LINK"}
	TypeAndThresh := map[string]float64{
		"VolChange24h": 100,
		"PerChange24h": 10,
		"PerChange1h":  5}

	From := "vietpride295@gmail.com"
	To := []string{"viettran295@gmail.com"}

	for {
		for _, symbol := range CryptoSym {
			go req.GetPrice(symbol, CryptoAPISrc, ch)
			payload := <-ch
			log.Println(payload)
			for typ, thresh := range TypeAndThresh {
				value := processor.ProcessCryptoAPIType(payload, symbol, typ)
				AbsVal := math.Abs(float64(value))

				if AbsVal > float64(thresh) {
					Subject, Msg := go_mail.CreateAlertMsg(symbol, typ, float64(value))
					log.Println("ALERT! ", Msg)
					go go_mail.SendEmail(From, To, Subject, Msg)
				}
			}
		}
		log.Println("Time execute: ", time.Since(start))
		time.Sleep(6 * time.Hour)
	}

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
