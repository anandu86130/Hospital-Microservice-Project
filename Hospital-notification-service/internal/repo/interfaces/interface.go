package interfaces

import "github.com/anandu86130/hospital-notification-service/internal/models"

type NotificationInter interface {
	NotificationStore(notify models.Notification) error
}
