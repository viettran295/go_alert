package chart_trend

import (
	"fmt"
	"go_alert/util"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type APISource struct {
	Url    string
	Method string
	ApiKey string
	Sym    string
}

func GetPrice() {
	cfg, _ := util.LoadConfig(".")

	apiSource := &APISource{
		Url:    "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest",
		Method: "GET",
		ApiKey: cfg.CoinMarketCapAPIkey,
		Sym:    "eth",
	}

	req, err := http.NewRequest("GET", apiSource.Url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("symbol", apiSource.Sym)
	req.Header.Add("X-CMC_PRO_API_KEY", apiSource.ApiKey)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	info, _ := io.ReadAll(resp.Body)
	fmt.Println(string(info))
}
