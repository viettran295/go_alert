package util

import "github.com/spf13/viper"

type Config struct{
	EmailSenderAddress string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error){
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	err = viper.ReadInConfig()
	if err != nil{
		return
	}
	err = viper.Unmarshal(&config)
	return
}