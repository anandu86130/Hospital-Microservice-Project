package repositories

import (
	"github.com/anandu86130/Hospital-admin-service/internal/model"
)

func (a *AdminRepository) FindAdminByEmail(email string) (*model.Admin, error) {
	var admin model.Admin
	if err := a.DB.Model(&model.Admin{}).Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
