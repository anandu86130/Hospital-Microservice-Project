package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Host         string `mapstructure:"HOST"`
	User         string `mapstructure:"USER"`
	Password     string `mapstructure:"PASSWORD"`
	Database     string `mapstructure:"DBNAME"`
	Port         string `mapstructure:"PORT"`
	Sslmode      string `mapstructure:"SSL"`
	GrpcPort     string `mapstructure:"GRPCPORT"`
	KAFKA_BROKER string `mapstructure:"KAFKA_BROKER"`
	APPPASSWORD  string `mapstructure:"APPPASSWORD"`
	APPEMAIL     string `mapstructure:"APPEMAIL"`
}

func LoadConfig() *Config {
	var config Config
	viper.SetConfigFile(".env")

	// Handle error if config file can't be read
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Handle error if config can't be unmarshaled
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to unmarshal config: %v", err)
	}

	// Debugging: Print the loaded config
	fmt.Printf("Loaded config: %+v\n", config)

	return &config
}
