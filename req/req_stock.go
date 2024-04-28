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

// @Params today: to query today stock price (true)
// 				  to query previous day stock price (false)
func requestAllStockPrice(today bool) (interface{} , error) {
	currTime := time.Now()
	currMonth := currTime.Month()
	currDay := currTime.Day()

	// Check whether today is weeken
	switch{
	case !today && int(currTime.Weekday()) == int(time.Saturday):
		currDay--
	case !today && int(currTime.Weekday()) == int(time.Sunday):
		return nil, errors.New("Fail to query data on weekend") 
	case !today && (currTime.Weekday()) == (time.Monday):
		currDay -= 2
	case today && (int(currTime.Weekday()) == int(time.Saturday) || int(currTime.Weekday()) == int(time.Sunday)):
		return nil, errors.New("Fail to query data on weekend") 
	}

	params := models.GetGroupedDailyAggsParams{
		Locale: "us",
		MarketType: "stocks",
		Date:   models.Date(time.Date(currTime.Year(), currMonth, currDay, 10, 0, 0, 0, time.UTC)),
	}.WithAdjusted(true)
	return c.GetGroupedDailyAggs(context.Background(), params)
	
}

// Get all stocks price and map ticker with price
func GetAllStockPrice(today bool) map[string]interface{} {
	result := make(map[string]interface{})
	resp, err := requestAllStockPrice(today)
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