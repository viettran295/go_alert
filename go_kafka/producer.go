package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)
func NewProducer(bootstrap_server string, group_id string) any {
	config := kafka.ConfigMap{
		"bootstrap.servers": bootstrap_server,
		"group.id":          group_id,
		"acks":              "all",
	}
	producer, err := kafka.NewProducer(&config)
	if err != nil {
		log.Println("Fail to creat producer")
	}
	return producer
}
func Kafka_produce(producer any, topic string, mess string) {
	delivery_report := make(chan kafka.Event, 10000)
	kafkaMess := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(mess),
	}
	err := producer.(*kafka.Producer).Produce(&kafkaMess, delivery_report)
	if err != nil {
		log.Println("Error in producing mess")
	}
	e := <-delivery_report
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil{
		log.Println("Delivery fail: ", m.TopicPartition.Error)
	}else{
		log.Printf("Delivered to topic: %s, partition: [%d] at offset: %v \n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		close(delivery_report)
	}

}
