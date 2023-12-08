package util

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	EmailSenderAddress  string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string `mapstructure:"EMAIL_SENDER_PASSWORD"`

	KafkaBootstrapServer string `mapstructure:"KAFKA_BOOTSTRAPSERVER"`
	ProducerGroupid      string `mapstructure:"PRODUCER_GROUP_Id"`
	ConsumerGroupid      string `mapstructure:"CONSUMER_GROUP_ID"`
	CoinMarketCapAPIkey  string `mapstructure:"COIN_MARKET_CAP_API_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	err = viper.ReadInConfig()
	if err != nil {
		log.Panicln("Error while loading config")
		return
	}
	err = viper.Unmarshal(&config)
	return
}
