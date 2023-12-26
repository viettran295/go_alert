package main

import (
	// kafka "go_alert/go_kafka"
	"go_alert/db"
	"go_alert/go_mail"
	"go_alert/processor"
	"go_alert/req"
	"log"
	"math"
	"time"
)

func main() {
	rdb := db.NewRdb()
	stockCh := make(chan req.StockResponse)
	ch := make(chan req.CryptoResponse)
	CryptoSym := []string{"BTC", "ETH", "SOL", "XRP", "LINK"}
	StockSym := []string{"GOOG", "COIN", "AMZN", "META", "MSTR"}

	TypeStockThresh := map[string]float64{
		"Market price": 10,
		"Volume":       50,
	}
	TypeAndThresh := map[string]float64{
		"VolChange24h": 100,
		"PerChange24h": 10,
		"PerChange1h":  5}

	for {
		for _, symbol := range StockSym {
			go req.GetStockPrice(symbol, stockCh)
			payload := <-stockCh
			log.Println(payload)

			newPrice := payload.Chart.Result[0].Meta.MarketPrice
			newVol := payload.Chart.Result[0].Indicators.Quote[0].Volume[0]
			oldPrice := db.GetRdb(&rdb, symbol)
			oldVol := db.GetRdb(&rdb, symbol)
			percentPriceChange := processor.PercentChange(oldPrice, newPrice)
			percentVolChange := processor.PercentChange(oldVol, float64(newVol))

			if percentPriceChange >= TypeStockThresh["Market price"] {
				go go_mail.CreateAlertMsg(symbol, "Market price", percentPriceChange)
			}
			if percentVolChange >= TypeStockThresh["Volume"] {
				go go_mail.CreateAlertMsg(symbol, "Volume", percentVolChange)
			}
			db.SetRdb(&rdb, symbol, newPrice)
			db.SetRdb(&rdb, symbol, newVol)

		}
		for _, symbol := range CryptoSym {
			go req.GetCryptoPrice(symbol, ch)
			payload := <-ch
			log.Println(payload)
			for typ, thresh := range TypeAndThresh {
				value := processor.ProcessCryptoType(payload, symbol, typ)
				AbsVal := math.Abs(float64(value))

				if AbsVal > float64(thresh) {
					go go_mail.CreateAlertMsg(symbol, typ, float64(value))
				}
			}
		}
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
