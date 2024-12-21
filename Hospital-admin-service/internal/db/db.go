package db

import (
	"fmt"
	"log"

	"github.com/anandu86130/Hospital-admin-service/config"
	"github.com/anandu86130/Hospital-admin-service/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *config.Config) *gorm.DB {
	host := config.Host
	user := config.User
	password := config.Password
	dbname := config.Database
	port := config.Port
	sslmode := config.Sslmode

	// Print each configuration for debugging
	fmt.Printf("Connecting to DB: host=%s, user=%s, password=%s, dbname=%s, port=%s, sslmode=%s\n", host, user, password, dbname, port, sslmode)

	// Construct the DSN correctly
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
	fmt.Println("DSN:", dsn) // Debugging line

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connection to the database failed:", err)
	}

	err = DB.AutoMigrate(
		&model.Admin{},
	)
	if err != nil {
		fmt.Printf("error while migrating %v", err.Error())
		return nil
	}

	// Check if an admin record already exists
	var existingAdmin model.Admin
	DB.First(&existingAdmin, "email = ?", config.Admin)
	if existingAdmin.ID == 0 { // No existing admin found
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(config.AdminPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal("failed to hash password:", err)
		}

		admin := model.Admin{
			Email:    config.Admin,
			Password: string(hashedPassword),
		}
		if err := DB.Create(&admin).Error; err != nil {
			log.Fatal("failed to create admin:", err)
		} else {
			fmt.Println("Admin user created successfully.")
		}
	} else {
		fmt.Println("Admin user already exists.")
	}
	return DB
}
