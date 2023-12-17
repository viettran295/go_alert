package req

import (
	"encoding/json"
	"go_alert/util"
	"io"
	"log"
	"net/http"
	"net/url"
)

type APISource struct {
	Url    string
	Method string
	ApiKey string
	Sym    string
}

type CryptoAPIResponse struct {
	Data map[string] struct {
		Quote struct {
			Usd CurrencyUSD `json:"USD"`
		} `json:"quote"`
	} `json:"data"`
}

type CurrencyUSD struct {
	Price            float32 `json:"price"`
	Volume24h        float32 `json:"volume_24h"`
	VolumeChange24h  float32 `json:"volume_change_24h"`
	PercentChange1h  float32 `json:"percent_change_1h"`
	PercentChange24h float32 `json:"percent_change_24h"`
	PercentChange60d float32 `json:"percent_change_60d"`
	PercentChange90d float32 `json:"percent_change_90d"`
	LastUpdate       string  `json:"last_updated"`
}

func processJSON(payload []byte) CryptoAPIResponse {
	resp := &CryptoAPIResponse{}
	if err := json.Unmarshal(payload, resp); err != nil {
		log.Panicln("Fail to process JSON")
	}
	return *resp
}

func GetPrice(symbol string, apiSrc APISource, ch chan CryptoAPIResponse){
	cfg, _ := util.LoadConfig(".")
	apiSrc.ApiKey = cfg.CoinMarketCapAPIkey
	apiSrc.Sym = symbol

	req, err := http.NewRequest("GET", apiSrc.Url, nil)
	if err != nil {
		log.Panicln("Fail to request")
	}

	q := url.Values{}
	q.Add("symbol", apiSrc.Sym)
	req.Header.Add("X-CMC_PRO_API_KEY", apiSrc.ApiKey)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	info, _ := io.ReadAll(resp.Body)
	ch <- processJSON(info)
}
