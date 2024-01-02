package req

import (
	"encoding/json"
	"fmt"
	"io"

	"log"
	"net/http"
)

type StockResponse struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Currency    string  `json:"currency"`
				Symbol      string  `json:"symbol"`
				MarketPrice float64 `json:"regularMarketPrice"`
			} `json:"meta"`
			Timestamp  []int32 `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Open   []float64 `json:"open"`
					Volume []int32   `json:"volume"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
	} `json:"chart"`
}

func processJSON[T any](payload []byte, resp *T) T {
	if err := json.Unmarshal(payload, resp); err != nil {
		log.Panicln("Fail to process JSON")
	}
	return *resp
}

func GetStockPrice(symbol string, ch chan StockResponse) {
	stockUrl := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?interval=1d", symbol)
	resp, err := http.Get(stockUrl)
	if err != nil {
		log.Panicln("Error while getting stock price")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("Error while reading stock price")
	}

	stockResp := processJSON(data, &StockResponse{})
	ch <- stockResp
}
