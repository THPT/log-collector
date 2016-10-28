package repo

import (
	"log-collector/app/model"
	"log-collector/infra"

	"github.com/Shopify/sarama"
)

type kafka struct{}

var Kafka kafka

func (kafka) SendEventLogAsync(event *model.Event, topic string, key string) {
	infra.Kafka.KafkaAsyncProducer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: event,
	}
}

func (kafka) SendEventLogSync(event *model.Event, topic string, key string) error {
	_, _, err := infra.Kafka.KafkaSyncProducer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: event,
	})
	return err
}
