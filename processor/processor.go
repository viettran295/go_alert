package processor

import (
	"go_alert/req"
	"log"
	"slices"
)

func convertRawToVolumeChange24h(payload req.CryptoAPIResponse, symbol string) float32 {
	return payload.Data[symbol].Quote.Usd.VolumeChange24h
}

func convertRawToPercentChange24h(payload req.CryptoAPIResponse, symbol string) float32 {
	return payload.Data[symbol].Quote.Usd.PercentChange24h
}

func convertRawToPercentChange1h(payload req.CryptoAPIResponse, symbol string) float32 {
	return payload.Data[symbol].Quote.Usd.PercentChange1h
}

func ProcessCryptoAPIType(payload req.CryptoAPIResponse, symbol string, typeComp string) float32 {
	supportedTypeComp := []string{"VolChange24h", "PerChange24h", "PerChange1h"}
	if slices.Contains(supportedTypeComp, typeComp) == false {
		log.Println("Type is not supported")
	}

	switch typeComp {
	case "VolChange24h":
		return convertRawToVolumeChange24h(payload, symbol)
	case "PerChange24h":
		return convertRawToPercentChange24h(payload, symbol)
	case "PerChange1h":
		return convertRawToPercentChange1h(payload, symbol)
	}
	return 0
}
