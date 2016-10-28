package infra

import (
	"crypto/tls"
	"fmt"
	"log-collector/config"
	"time"

	"github.com/Shopify/sarama"
)

type kafka struct {
	KafkaAsyncProducer sarama.AsyncProducer
	KafkaSyncProducer  sarama.SyncProducer
	KafkaConsumer      sarama.Consumer
}

var (
	Kafka kafka
)

func CloseKafka() {
	fmt.Println("Close all connect resource...")
	Kafka.KafkaAsyncProducer.Close()
	Kafka.KafkaSyncProducer.Close()
	Kafka.KafkaConsumer.Close()
}

func initKafkaAsyncProducer() (sarama.AsyncProducer, error) {
	conf := sarama.NewConfig()
	tlsConfig := createTlsConfiguration()
	if tlsConfig != nil {
		conf.Net.TLS.Enable = true
		conf.Net.TLS.Config = tlsConfig
	}
	conf.Producer.RequiredAcks = sarama.WaitForLocal
	conf.Producer.Compression = sarama.CompressionSnappy
	conf.Producer.Flush.Frequency = 500 * time.Millisecond
	producer, err := sarama.NewAsyncProducer([]string{config.KafkaHost + ":" + config.KafkaPort}, conf)
	return producer, err
}

func initKafkaSyncProducer() (sarama.SyncProducer, error) {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Retry.Max = 10
	tlsConfig := createTlsConfiguration()
	if tlsConfig != nil {
		conf.Net.TLS.Config = tlsConfig
		conf.Net.TLS.Enable = true
	}
	producer, err := sarama.NewSyncProducer([]string{config.KafkaHost + ":" + config.KafkaPort}, conf)
	return producer, err
}

func initKafkaConsumer() (sarama.Consumer, error) {
	conf := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{config.KafkaHost + ":" + config.KafkaPort}, conf)
	return consumer, err
}

func createTlsConfiguration() (t *tls.Config) {
	return nil
}

func InitKafka() {
	asyncProducer, err := initKafkaAsyncProducer()
	if err != nil {
		fmt.Println("Connect Kafka failed...", err)
		panic(err)
	}
	Kafka.KafkaAsyncProducer = asyncProducer

	syncProducer, err := initKafkaSyncProducer()
	if err != nil {
		fmt.Println("Connect Kafka failed...", err)
		panic(err)
	}
	Kafka.KafkaSyncProducer = syncProducer

	kafkaConsumer, err := initKafkaConsumer()
	if err != nil {
		fmt.Println("Connect Kafka consumer failed...", err)
		panic(err)
	}
	Kafka.KafkaConsumer = kafkaConsumer

}
