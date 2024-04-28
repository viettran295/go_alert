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
	stockCurrCh := make(chan map[string]float64)
	stockPrevCh := make(chan map[string]float64)

	StockThresh := 5
	TypeAndThresh := map[string]float64{
		"VolChange24h": 100,
		"PerChange24h": 10,
		"PerChange1h":  5,
	}

	for {
		go req.GetAllStockPrice(true, stockCurrCh)
		go req.GetAllStockPrice(false, stockPrevCh)
		select{
		case todayStockPrice := <- stockCurrCh:
			select{
			case prevStockPrice := <- stockPrevCh: 
				for _, ticker := range StockSym {
					log.Println("Today Stock: ", ticker, "Price: ",todayStockPrice[ticker])
					log.Println("Previous Stock: ", ticker, "Price: ",prevStockPrice[ticker])
					percentChange := processor.PercentChange(todayStockPrice[ticker], prevStockPrice[ticker]) 
					if percentChange > float64(StockThresh){
						log.Printf("ALERT percent price change of %s: %f", ticker, percentChange)
						go go_mail.CreateAlertMsg(ticker , "Percent price", percentChange)
					}
				}
			default:
			}
		default:
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
		time.Sleep(12 * time.Hour)
	}

	// <<<<<< Kafka >>>>>>
	// if err_cfg != nil {
	// 	log.Println("Can not load config info")
	// }
	// consumer := kafka.NewConsumer(cfg.KafkaBootstrapServer, cfg.ConsumerGroupid)
	// kafka.Kafka_consume(consumer, []string{"saigon"})
}