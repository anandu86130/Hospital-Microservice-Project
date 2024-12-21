package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/anandu86130/hospital-notification-service/internal/models"
)

// SubscribeAndConsumeCuttingEvents implements interfaces.NotificationServiceInter.
func (n *NotificationService) SubScribeAndConsumeAppointmentEvents() error {
	for {
		// cuttingConsumer is similar to paymentConsumer
		msg, err := n.AppointmentResultConsumer.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Consumer error", err)
			continue
		}

		var appointmentevent models.AppointmentResultEvent
		if err := json.Unmarshal(msg.Value, &appointmentevent); err != nil {
			fmt.Printf("Failed to unmarshal cutting event message: %v\n", err)
			continue
		}
		fmt.Printf("Raw Message as String: %s\n", string(msg.Value))

		fmt.Printf("Received appointment Event: %v\n", appointmentevent)
	}
}
