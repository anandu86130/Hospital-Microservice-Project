package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/anandu86130/hospital-notification-service/internal/models"
	utils "github.com/anandu86130/hospital-notification-service/internal/utils"
)

// SubscribeAndConsumePaymentEvents implements interfaces.NotificationServiceInter.
func (n *NotificationService) SubscribeAndConsumePaymentEvents() error {
	log.Println("waiting for the event")
	for {
		msg, err := n.paymentConsumer.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Consumer error", err)
			continue
		} else {
			fmt.Println("Message received: ", msg)
		}

		var paymentEvent models.PaymentEvent
		if err := json.Unmarshal(msg.Value, &paymentEvent); err != nil {
			fmt.Printf("Failed to unmarshal payment message: %v\n", err)
			continue
		}
		fmt.Println("receiving produce", paymentEvent)

		err = utils.SendNotificationToEmail(paymentEvent, "Payment Notification", fmt.Sprintf("Payment received: %.d", paymentEvent.Amount))

		if err != nil {
			fmt.Printf("Error sending notification: %v\n", err)
		} else {
			fmt.Println("Notification sent successfully!")
		}
	}
}
