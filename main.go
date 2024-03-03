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
	stockCh := make(chan req.StockResponse)
	ch := make(chan req.CryptoResponse)
	CryptoSym := []string{"BTC", "ETH", "SOL", "XRP", "LINK"}
	StockSym := []string{"GOOG", "COIN", "AMZN", "META", "MSTR",
		"AMD", "ARM", "NVDA", "TXN", "IBM"}

	TypeStockThresh := map[string]float64{
		"Market price": 10,
		"Volume":       20,
	}
	TypeAndThresh := map[string]float64{
		"VolChange24h": 100,
		"PerChange24h": 10,
		"PerChange1h":  5,
	}

	for {
		for _, symbol := range StockSym {
			go req.GetStockPrice(symbol, stockCh)
			payload := <-stockCh
			log.Println(payload)
			highPrice := payload.Chart.Result[0].Indicators.Quote[0].High[0]
			lowPrice := payload.Chart.Result[0].Indicators.Quote[0].Low[0]

			percentChange := processor.PercentChange(highPrice, lowPrice)
			log.Println("Percent change of price: ", percentChange)

			if percentChange >= TypeStockThresh["Market price"] {
				log.Printf("ALERT percent price change of %s: %f \n", symbol, percentChange)
				go go_mail.CreateAlertMsg(symbol, "Percent price change", percentChange)
			}
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
