package db

import (
	"fmt"
	"log"

	"github.com/anandu86130/hospital-notification-service/config"
	"github.com/anandu86130/hospital-notification-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *config.Config) *gorm.DB {
	// Validate the configuration before constructing the DSN
	if config.Host == "" || config.User == "" || config.Password == "" || config.Database == "" || config.Port == "" || config.Sslmode == "" {
		log.Fatal("missing required database configuration values")
	}

	// Construct the DSN correctly
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host, config.User, config.Password, config.Database, config.Port, config.Sslmode,
	)

	// Debugging: Print the DSN (without sensitive info)
	fmt.Printf("Connecting to DB: host=%s, user=%s, dbname=%s, port=%s, sslmode=%s\n", config.Host, config.User, config.Database, config.Port, config.Sslmode)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connection to the database failed: %v", err)
	}

	err = DB.AutoMigrate(
		&models.Notification{},
		&models.PaymentEvent{},
	)
	if err != nil {
		log.Fatalf("error when migrating: %v", err)
	}

	log.Println("Database connection established successfully")
	return DB
}
