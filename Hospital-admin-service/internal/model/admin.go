package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Email    string
	Password string
}