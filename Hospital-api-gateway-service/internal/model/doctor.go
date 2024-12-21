package model

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	Name              string `json:"name" gorm:"not null"`
	Email             string `json:"email" gorm:"not null; unique"`
	Password          string `json:"password" gorm:"not null"`
	Specialization    string `json:"specialization" gorm:"not null"`
	YearsOfExperience uint32 `json:"years_of_experience" gorm:"not null"`
	Fees              uint32 `json:"fees" gorm:"not null"`
}

type EditDoctor struct {
	gorm.Model
	Name              string `json:"name" gorm:"not null"`
	YearsOfExperience uint32 `json:"years_of_experience" gorm:"not null"`
	Fees              uint32 `json:"fees" gorm:"not null"`
}

type OTP struct {
	Email    string `json:"email"`
	Otp      string `json:"otp"`
	Password string `json:"password"`
}
