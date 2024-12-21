package utility

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type AppointmentResultEvent struct {
	AppointmentID uint                 `json:"appointment_id"`
	Appointments  []AppointmentPayload `json:"appointments"`
}

type AppointmentPayload struct {
	AppointmentID uint32    `json:"appointment_id"`
	Doctorid      uint32    `gorm:"column:doctorid" json:"doctorid"`
	Userid        uint32    `gorm:"column:userid" json:"userid"`
	Date          time.Time `gorm:"type:date" json:"date"`
	StartTime     time.Time `gorm:"type:time" json:"starttime"`
	EndTime       time.Time `gorm:"type:time" json:"endtime"`
}

type KafkaAppointmentResultProducer struct {
	writer *kafka.Writer
}

func NewKafkaAppointmentResultProducer(broker string) (*KafkaAppointmentResultProducer, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{broker},
		Topic:        "appointment_topic",
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: int(kafka.RequireOne),
	})
	return &KafkaAppointmentResultProducer{writer: writer}, nil
}

func (k *KafkaAppointmentResultProducer) ProducerAppointmentResultEvent(event AppointmentResultEvent) error {
	message, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}
	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("%d", event.AppointmentID)),
		Value: message,
	}
	err = k.writer.WriteMessages(context.Background(), msg)
	if err != nil {
		return fmt.Errorf("failed to produce appointment result event: %w", err)
	}
	log.Printf("appointment result event successfully sent to Kafka topic: %s", k.writer.Topic)
	return nil
}

func HandleAppointmentResultNotification(appointmentID uint32, appointmnets []AppointmentPayload) error {
	kafkaProducer, err := NewKafkaAppointmentResultProducer("kafka:9092")
	if err != nil {
		return fmt.Errorf("failed to create Kafka producer: %w", err)
	}

	// Create and send CuttingResultEvent
	appointmentresultevent := AppointmentResultEvent{
		Appointments: appointmnets,
	}
	err = kafkaProducer.ProducerAppointmentResultEvent(appointmentresultevent)
	if err != nil {
		return fmt.Errorf("failed to produce appointment result event: %w", err)
	}

	log.Println("appointment result event produced successfully")
	return nil
}
