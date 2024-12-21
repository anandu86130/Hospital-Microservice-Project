package handler

import "github.com/anandu86130/hospital-notification-service/internal/services/interfaces"

type notifcationHandler struct {
	services interfaces.NotificationServiceInter
}

func NewNotificationHandler(service interfaces.NotificationServiceInter) *notifcationHandler {
	return &notifcationHandler{
		services: service,
	}
}
