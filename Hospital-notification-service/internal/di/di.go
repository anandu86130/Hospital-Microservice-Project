package di

import (
	"log"
	"sync"

	"github.com/anandu86130/hospital-notification-service/config"
	"github.com/anandu86130/hospital-notification-service/internal/db"
	"github.com/anandu86130/hospital-notification-service/internal/handler"
	"github.com/anandu86130/hospital-notification-service/internal/kafka"
	"github.com/anandu86130/hospital-notification-service/internal/repo"
	"github.com/anandu86130/hospital-notification-service/internal/services"
)

func Init() {
	cfg := config.LoadConfig()

	dbconn := db.ConnectDB(cfg)

	// Initialize Kafka consumers
	paymentConsum, err := kafka.NewKafkaConsumer(cfg.KAFKA_BROKER, "payment_service_group", "payment_topic")
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer for payment topic: %v", err)
	}
	appointmentConsum, err := kafka.NewKafkaConsumer(cfg.KAFKA_BROKER, "cutting_service_group", "cutting_topic")
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer for cutting topic: %v", err)
	}

	repo := repo.NewRepository(dbconn)
	// Initialize the NotificationService
	srv := services.NewNotificationService(repo, paymentConsum, appointmentConsum)
	handl := handler.NewNotificationHandler(srv)

	var wg sync.WaitGroup

	// Start payment handler in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := handl.PaymentHandler()
		if err != nil {
			log.Fatalf("Error in payment consumer: %v", err)
		}
	}()

	// Start cutting result handler in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := handl.AppointmentResultHandler()
		if err != nil {
			log.Fatalf("Error in cutting consumer: %v", err)
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
}
