package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewConsumer(bootstrap_server string, group_id string) any {
	config := &kafka.ConfigMap{
		"bootstrap.servers": bootstrap_server,
		"group.id":          group_id,
		"auto.offset.reset": "smallest",
	}
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Println("Error in init Consumer")
	}
	return consumer
}

func Kafka_consume(consumer any, topics []string) {
	c := consumer.(*kafka.Consumer)
	err := c.SubscribeTopics(topics, nil)
	if err != nil {
		log.Println("Error in subcribtion topics")
	}
	for {
		ev := c.Poll(1000)
		switch e := ev.(type) {
		case *kafka.Message:
			log.Println("Message: ", string(e.Value))
		case kafka.Error:
			log.Println(e)
			break
		default:
			log.Println("Ignore: ", e)
		}
	}
	c.Close()
}
