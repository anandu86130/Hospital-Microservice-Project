package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/anandu86130/Hospital-doctor-service/internal/model"
	inter "github.com/anandu86130/Hospital-doctor-service/internal/repository/interfaces"
	"gorm.io/gorm"
)

type DoctorRepository struct {
	DB *gorm.DB
}

// CheckDoctorExistsByLiscenceNumber implements interfaces.DoctorRepository
func NewDoctorRepository(db *gorm.DB) inter.DoctorRepository {
	return &DoctorRepository{
		DB: db,
	}
}

func (d *DoctorRepository) CreateDoctor(doctor *model.Doctor) error {
	if err := d.DB.Create(&doctor).Error; err != nil {
		return err
	}
	return nil
}

func (d *DoctorRepository) CreateOTP(otp *model.OTP) error {
	if err := d.DB.Create(&otp).Error; err != nil {
		return err
	}
	return nil
}

func (d *DoctorRepository) FindDoctorByEmail(email string) (*model.Doctor, error) {
	var doctor model.Doctor
	if email == "" {
		return &doctor, fmt.Errorf("email cannot be empty")
	}

	log.Printf("Querying for user with email: %s\n", email)

	err := d.DB.Where("email = ?", email).First(&doctor).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &doctor, fmt.Errorf("user with email %s not found", email)
	} else if err != nil {
		return &doctor, fmt.Errorf("failed to find user by email: %w", err)
	}
	return &doctor, nil
}

func (d *DoctorRepository) VerifyOTPcheck(email string, otp string) error {
	var doctor model.OTP
	result := d.DB.Where("email = ? AND otp = ?", email, otp).Find(&doctor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("email or OTP not found")
		}
		return fmt.Errorf("failed to verify OTP: %w", result.Error)
	}
	return nil
}

func (d *DoctorRepository) FindOTPByEmail(email string) (*model.OTP, error) {
	var otp model.OTP
	err := d.DB.Where("email = ?", email).First(&otp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // No OTP found for the given email
		}
		return nil, err // Other errors
	}
	return &otp, nil
}

func (d *DoctorRepository) UpdateOTP(otp *model.OTP) error {
	err := d.DB.Model(&model.OTP{}).Where("email = ?", otp.Email).Update("otp", otp.Otp).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *DoctorRepository) DoctorExists(email string) (bool, error) {
	var count int64
	err := d.DB.Model(&model.Doctor{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (d *DoctorRepository) UpdateDoctor(doctor *model.Doctor) error {
	// Use Updates method to update only specific fields
	return d.DB.Model(&model.Doctor{}).Where("id = ?", doctor.ID).Updates(map[string]interface{}{
		"name":                doctor.Name,
		"years_of_experience": doctor.YearsOfExperience,
		"fees":                doctor.Fees,
	}).Error
}

func (d *DoctorRepository) GetDoctorList() ([]*model.Doctor, error) {
	var doctor []*model.Doctor
	if err := d.DB.Find(&doctor).Error; err != nil {
		return nil, err
	}
	return doctor, nil
}

func (d *DoctorRepository) FindDoctorByID(doctorID uint) (*model.Doctor, error) {
	var doctor model.Doctor
	if err := d.DB.First(&doctor, doctorID).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (d *DoctorRepository) DoctorDetails(doctorID uint) (*model.Doctor, error) {
	var doctor model.Doctor
	result := d.DB.Where("id = ? ", doctorID).First(&doctor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &doctor, fmt.Errorf("email or OTP not found")
		}
		return &doctor, fmt.Errorf("failed to verify OTP: %w", result.Error)
	}
	return &doctor, nil
}
