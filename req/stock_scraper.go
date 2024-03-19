package req

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Stock struct {
	Company       string
	Price, Change float64
}

func trimSpecialChar(str string) string {
	// Extract number from web e.g: (+3.5%) or 1,700.55
	number := regexp.MustCompile(`[-+]?\d+([\,]\d+)*([\.]\d+)`)
	return number.FindString(str)
}

// Covert price and percent change from string to float
// to match with Stock type attribute
func toStockType(str string) float64 {
	number := trimSpecialChar(str)
	toFloat, err := strconv.ParseFloat(strings.ReplaceAll(number, ",", ""), 32)
	if err != nil {
		log.Println("Error while convert string to float: ", err)
	}

	return toFloat
}

func ScrapStock(ticker string, ch chan Stock) {
	stock := Stock{}
	url := "https://finance.yahoo.com/quote/" + ticker

	c := colly.NewCollector()
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error while scrapping: ", err)
	})
	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {
		stock.Company = e.ChildText("h1")
		price := e.ChildText("fin-streamer[data-field='regularMarketPrice']")
		stock.Price = toStockType(price)
		change := e.ChildText("fin-streamer[data-field='regularMarketChangePercent']")
		stock.Change = toStockType(change)
		ch <- stock
	})
	go c.Visit(url)
}
