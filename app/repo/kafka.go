package repo

import (
	"fmt"
	"log-collector/app/model"
	"log-collector/infra"

	"github.com/Shopify/sarama"
)

type kafka struct{}

var Kafka kafka

func (kafka) SendEventLogAsync(event *model.Event, topic string, key string) {
	fmt.Println(topic)
	infra.Kafka.KafkaAsyncProducer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: event,
	}
}
