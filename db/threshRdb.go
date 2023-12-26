package db

import (
	"log"
	"strconv"

	"github.com/go-redis/redis"
)

func NewRdb() redis.Client {
	config := redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	client := redis.NewClient(&config)
	return *client
}

func SetRdb(rdb any, key string, value any) {
	err := rdb.(*redis.Client).Set(key, value, 0).Err()
	if err != nil {
		log.Panicln("Error while setting value to redis")
	}
}

func GetRdb(rdb *redis.Client, key string) float64 {
	val, _ := rdb.Get(key).Result()
	ret, _ := strconv.ParseFloat(val, 64)
	return ret
}
