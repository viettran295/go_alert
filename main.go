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
	CryptoSym := []string{"BTC", "ETH", "SOL", "XRP", "LINK"}
	StockSym := []string{"GOOG", "COIN", "AMZN", "META", "MSTR",
		"AMD", "ARM", "NVDA", "TXN", "IBM"}

	ch := make(chan req.CryptoResponse, len(CryptoSym))
	stockCh := make(chan req.Stock, len(StockSym))

	StockThresh := map[string]float32{
		"PriceThresh": 5,
	}

	TypeAndThresh := map[string]float64{
		"VolChange24h": 100,
		"PerChange24h": 10,
		"PerChange1h":  5,
	}

	for {
		for _, ticker := range StockSym {
			go req.ScrapStock(ticker, stockCh)
		}
		for i := 0; i < len(StockSym); i++ {
			stock := <-stockCh
			log.Println(stock)
			if stock.Change > StockThresh["PriceThresh"] {
				log.Printf("ALERT percent price change of %s: %f", stock.Company, stock.Change)
				go go_mail.CreateAlertMsg(stock.Company, "Percent price", float64(stock.Change))
			}
		}

		for _, symbol := range CryptoSym {
			go req.GetCryptoPrice(symbol, ch)
		}
		for i := 0; i < len(CryptoSym); i++ {
			payload := <-ch
			log.Print(payload)
			for typ, thresh := range TypeAndThresh {
				value := processor.ProcessCryptoType(payload, CryptoSym[i], typ)
				AbsVal := math.Abs(float64(value))

				if AbsVal > float64(thresh) {
					log.Printf("ALERT percent price change of %s: %f \n", CryptoSym[i], value)
					go go_mail.CreateAlertMsg(CryptoSym[i], typ, float64(value))
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
