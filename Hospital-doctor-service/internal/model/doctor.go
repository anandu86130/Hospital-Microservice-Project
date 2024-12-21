package model

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	Name              string `json:"name" gorm:"not null"`
	Email             string `json:"email" gorm:"not null; unique"`
	Password          string `json:"password" gorm:"not null"`
	Specialization    string `json:"specialization" gorm:"not null"`
	YearsOfExperience uint32 `json:"years_of_experience" gorm:"not null"`
	LicenceNumber     string `json:"licence_number" gorm:"not null; unique"`
	Fees              uint32 `json:"fees" gorm:"not null"`
	IsBlocked         bool   `json:"is_blocked" gorm:"default:false"`
	IsVerified        bool   `json:"is_verified" gorm:"default:false"`
}

type Review struct {
	ID        uint `gorm:"id"`
	DoctorID  uint `gorm:"doctor_id"`
	PatientID uint `gorm:"patient_id"`
	Rating    uint `gorm:"rating"`
}

type OTP struct {
	Email    string `json:"email"`
	Otp      string `json:"otp"`
	Password string `json:"password"`
}
