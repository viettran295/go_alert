package util

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	EmailSenderAddress  string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string `mapstructure:"EMAIL_SENDER_PASSWORD"`

	KafkaBootstrapServer string `mapstructure:"KAFKA_BOOTSTRAPSERVER"`
	ProducerGroupid      string `mapstructure:"PRODUCER_GROUP_Id"`
	ConsumerGroupid      string `mapstructure:"CONSUMER_GROUP_ID"`
	CoinMarketCapAPIkey  string `mapstructure:"COIN_MARKET_CAP_API_KEY"`
	PolygonKey           string `mapstructure:"POLYGON_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	// Add credentials by .env file or pass as argument
	if _, err := os.Stat("app.env"); err == nil {
		viper.AddConfigPath(path)
		viper.SetConfigName("app")
		viper.SetConfigType("env")
		err = viper.ReadInConfig()
		if err != nil {
			log.Println("Error while loading credential in .env file")
		}
	} else{
		config.EmailSenderAddress = os.Getenv("EMAIL_SENDER_ADDRESS")
		config.EmailSenderPassword = os.Getenv("EMAIL_SENDER_PASSWORD")
		config.CoinMarketCapAPIkey = os.Getenv("COIN_MARKET_CAP_API_KEY")
		config.PolygonKey = os.Getenv("POLYGON_KEY")
	}
	err = viper.Unmarshal(&config)
	return
}
