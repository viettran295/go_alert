package processor

import (
	"go_alert/req"
	"log"
	"slices"
)

func convertRawToVolumeChange24h(payload req.CryptoResponse, symbol string) float32 {
	return payload.Data[symbol].Quote.Usd.VolumeChange24h
}

func convertRawToPercentChange24h(payload req.CryptoResponse, symbol string) float32 {
	return payload.Data[symbol].Quote.Usd.PercentChange24h
}

func convertRawToPercentChange1h(payload req.CryptoResponse, symbol string) float32 {
	return payload.Data[symbol].Quote.Usd.PercentChange1h
}

func ProcessCryptoType(payload req.CryptoResponse, symbol string, typeComp string) float32 {
	supportedTypeComp := []string{"VolChange24h", "PerChange24h", "PerChange1h"}
	if slices.Contains(supportedTypeComp, typeComp) == false {
		log.Panicln("Type is not supported")
	}

	switch typeComp {
	case supportedTypeComp[0]:
		return convertRawToVolumeChange24h(payload, symbol)
	case supportedTypeComp[1]:
		return convertRawToPercentChange24h(payload, symbol)
	case supportedTypeComp[2]:
		return convertRawToPercentChange1h(payload, symbol)
	default:
		return 0
	}
}
