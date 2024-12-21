package config

import "github.com/spf13/viper"

type Config struct {
	SECRETKEY       string `mapstructure:"JWTKEY"`
	Host            string `mapstructure:"HOST"`
	User            string `mapstructure:"USER"`
	Password        string `mapstructure:"PASSWORD"`
	Database        string `mapstructure:"DBNAME"`
	Port            string `mapstructure:"PORT"`
	Sslmode         string `mapstructure:"SSL"`
	GrpcDoctorPort  string `mapstructure:"GRPCDOCTORPORT"`
	RedisHost       string `mapstructure:"REDISHOST"`
	GrpcBookingPort string `mapstructure:"GRPCBOOKINGPORT"`
	GrpcUserPort    string `mapstructure:"GRPCUSERPORT"`
}

// LoadConfig will load the environment variable to access.
func LoadConfig() *Config {
	var config Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	return &config
}
