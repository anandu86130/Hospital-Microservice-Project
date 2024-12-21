package repo

import (
	"errors"

	"github.com/anandu86130/hospital-notification-service/internal/models"
)

// NotificationStore implements interfaces.NotificationInter.
func (m *Repository) NotificationStore(notify models.Notification) error {
	err := m.DB.Create(notify)
	if err != nil {
		return errors.New("error while creating notification")
	}

	return nil
}
