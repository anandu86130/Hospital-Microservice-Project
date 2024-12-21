package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config represents the environment variables.
type Config struct {
	APIPORT         string `mapstructure:"APIPORT"`
	SECRETKEY       string `mapstructure:"JWTKEY"`
	SECRETKEYDOCTOR string `mapstructure:"JWTKEYDOCTOR"`
	DOCTORPORT      string `mapstructure:"DOCTORPORT"`
	USERPORT        string `mapstructure:"USERPORT"`
	ADMINPORT       string `mapstructure:"ADMINPORT"`
	BOOKINGPORT     string `mapstructure:"BOOKINGPORT"`
	STRIPEKEY       string `mapstructure:"STRIPE_SECRET_KEY"`
	ChatPort        string `mapstructure:"CHATPORT"`
	KafkaPort       string `mapstructure:"KAFKA_BROKER"`
	KafkaTpic       string `mapstructure:"KAFKA_TOPIC"`
}

func LoadConfig() (*Config, error) {
	var config Config
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read configuration file: %w", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	// Validate critical configuration keys
	if config.SECRETKEY == "" {
		return nil, fmt.Errorf("JWTKEY is missing in configuration")
	}

	viper.SetConfigFile("/app/.env")

	return &config, nil
}
