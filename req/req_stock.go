package req

import (
	"context"
	"go_alert/util"
	"log"
	"time"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

var cfg, _ = util.LoadConfig(".")
var c = polygon.New(cfg.PolygonKey)

func requestAllStockPrice() (interface{} , error) {
	currTime := time.Now()
	currMonth := currTime.Month()
	params := models.GetGroupedDailyAggsParams{
		Locale: "us",
		MarketType: "stocks",
		Date:   models.Date(time.Date(currTime.Year(), currMonth, currTime.Day()-1, 10, 0, 0, 0, time.UTC)),
	}.WithAdjusted(true)
	return c.GetGroupedDailyAggs(context.Background(), params)
	
}

// Get all stocks price and map ticker with price
func GetAllStockPrice() map[string]interface{} {
	result := make(map[string]interface{})
	resp, err := requestAllStockPrice()
	if err != nil{
		log.Println(err)
		return nil
	} 
	stocks := resp.(*models.GetGroupedDailyAggsResponse).Results
	for _, stock := range(stocks){
		result[stock.Ticker] = stock.High
	}
	return result
}

// Get stock price in previous day
func GetPrevStockPrice(ticker string, ch chan float64) {
	params := models.GetPreviousCloseAggParams{
		Ticker: ticker,
	}.WithAdjusted(true)

	resp, err := c.GetPreviousCloseAgg(context.Background(), params)
	if err != nil {
		log.Println(err)
		log.Panicln("Fail to request stock price")
	}
	ch <- resp.Results[0].Low
}