package req

import (
	"context"
	"errors"
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
	if (currTime.Weekday() == time.Saturday) || (currTime.Weekday() == time.Sunday){
		return nil, errors.New("Fail to query data on weekend")
	}
	currMonth := currTime.Month()
	currDay := currTime.Day()
	
	params := models.GetGroupedDailyAggsParams{
		Locale: "us",
		MarketType: "stocks",
		Date:   models.Date(time.Date(currTime.Year(), currMonth, currDay, 10, 0, 0, 0, time.UTC)),
	}.WithAdjusted(true)
	return c.GetGroupedDailyAggs(context.Background(), params)
}

// Get all stocks price and map ticker with price
func GetAllStockPrice(ch chan map[string]float64) {
	result := make(map[string]float64)
	resp, err := requestAllStockPrice()
	if err == nil{
		stocks := resp.(*models.GetGroupedDailyAggsResponse).Results
		for _, stock := range(stocks){
			result[stock.Ticker + " high"] = stock.High
			result[stock.Ticker + " low"] = stock.Low
		}
		ch <-result
	} else {
		log.Println(err)
	}
}