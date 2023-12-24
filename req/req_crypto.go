package req

import (
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

type CryptoResponse struct {
	Data map[string] struct {
		Quote struct {
			Usd struct {
				Price            float32 `json:"price"`
				Volume24h        float32 `json:"volume_24h"`
				VolumeChange24h  float32 `json:"volume_change_24h"`
				PercentChange1h  float32 `json:"percent_change_1h"`
				PercentChange24h float32 `json:"percent_change_24h"`
				PercentChange60d float32 `json:"percent_change_60d"`
				PercentChange90d float32 `json:"percent_change_90d"`
				LastUpdate       string  `json:"last_updated"`
			} `json:"USD"`
		} `json:"quote"`
	} `json:"data"`
}

func GetCryptoPrice(symbol string, ch chan CryptoResponse){
	var apiSrc = APISource{
	Url:    "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest",
	Method: "GET",
	}

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
	ch <- processJSON(info, &CryptoResponse{})
}
