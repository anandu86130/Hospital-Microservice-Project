package repo

import (
	"github.com/anandu86130/hospital-notification-service/internal/repo/interfaces"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) interfaces.NotificationInter {
	return &Repository{
		DB: db,
	}
}
