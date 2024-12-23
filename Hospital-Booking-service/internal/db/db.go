package db

import (
	"booking-service/config"
	"booking-service/internal/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *config.Config) *gorm.DB {
	host := config.Host
	user := config.User
	password := config.Password
	dbname := config.Database
	port := config.Port
	sslMode := config.Sslmode

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port =%s sslmode=%s", host, user, password, dbname, port, sslMode)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connection to the database failed:", err)
	}
	err = DB.AutoMigrate(
		&model.Availability{},
		&model.Appoinment{},
		&model.Prescription{},
		&model.Payment{},
	)

	if err != nil {
		fmt.Printf("error in while migerating DB%v", err.Error())
		return nil
	}

	return DB
}
