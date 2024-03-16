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

func main() {
	ch := make(chan req.CryptoResponse)
	stockCh := make(chan req.Stock, 10)
	stock := req.Stock{}

	CryptoSym := []string{"BTC", "ETH", "SOL", "XRP", "LINK"}
	StockSym := []string{"GOOG", "COIN", "AMZN", "META", "MSTR",
		"AMD", "ARM", "NVDA", "TXN", "IBM"}

	TypeAndThresh := map[string]float64{
		"VolChange24h": 100,
		"PerChange24h": 10,
		"PerChange1h":  5,
	}

	for {
		for _, ticker := range StockSym {
			go req.ScrapStock(ticker, stockCh)
			stock = <-stockCh
			log.Println(stock)
		}

		for _, symbol := range CryptoSym {
			go req.GetCryptoPrice(symbol, ch)
			payload := <-ch
			log.Println(payload)
			for typ, thresh := range TypeAndThresh {
				value := processor.ProcessCryptoType(payload, symbol, typ)
				AbsVal := math.Abs(float64(value))

				if AbsVal > float64(thresh) {
					log.Printf("ALERT percent price change of %s: %f \n", symbol, value)
					go go_mail.CreateAlertMsg(symbol, typ, float64(value))
				}
			}
		}
		time.Sleep(6 * time.Hour)
	}

	// <<<<<< Kafka >>>>>>
	// if err_cfg != nil {
	// 	log.Println("Can not load config info")
	// }
	// consumer := kafka.NewConsumer(cfg.KafkaBootstrapServer, cfg.ConsumerGroupid)
	// kafka.Kafka_consume(consumer, []string{"saigon"})
}
