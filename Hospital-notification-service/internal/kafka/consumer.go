package kafka

import (
	"time"
	"github.com/segmentio/kafka-go"
)

func NewKafkaConsumer(broker, groupID, topic string) (*kafka.Reader, error) {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{broker},
		GroupID:     groupID,
		Topic:       topic,
		StartOffset: kafka.FirstOffset,
		MinBytes:    10e3,
		MaxBytes:    10e6,
		MaxWait:     1 * time.Second,
	})

	return reader, nil
}
