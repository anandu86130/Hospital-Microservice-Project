package db

import (
	"fmt"
	"log"

	"github.com/anandu86130/Hospital-doctor-service/config"
	"github.com/anandu86130/Hospital-doctor-service/internal/model"
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
		&model.Doctor{},
		&model.OTP{},
	)

	if err != nil {
		fmt.Printf("error in while migerating DB%v", err.Error())
		return nil
	}

	return DB
}
