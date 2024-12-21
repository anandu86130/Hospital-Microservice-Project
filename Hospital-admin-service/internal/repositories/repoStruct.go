package repositories

import (
	interfaces "github.com/anandu86130/Hospital-admin-service/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepoInter {
	return &AdminRepository{
		DB: db,
	}
}
