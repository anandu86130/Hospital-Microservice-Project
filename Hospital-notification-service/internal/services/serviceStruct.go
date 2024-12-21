package services

import (
	"github.com/anandu86130/hospital-notification-service/internal/repo/interfaces"
	inter "github.com/anandu86130/hospital-notification-service/internal/services/interfaces"
	"github.com/segmentio/kafka-go"
)

type NotificationService struct {
	Repo                      interfaces.NotificationInter
	paymentConsumer           *kafka.Reader
	AppointmentResultConsumer *kafka.Reader
}
func NewNotificationService(repo interfaces.NotificationInter, paymentCon, AppointmentResCon *kafka.Reader) inter.NotificationServiceInter {
	return &NotificationService{
		Repo:                      repo,
		paymentConsumer:           paymentCon,
		AppointmentResultConsumer: AppointmentResCon,
	}
}
