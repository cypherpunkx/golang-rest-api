package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
}

type Server struct {
	Port int `mapstructure:"SERVER_PORT"`
}

type Database struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	Username string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	DBName   string `mapstructure:"DB_NAME"`
}

// LoadConfig membaca dan mengembalikan konfigurasi dari berkas .env
func LoadConfig(configFile string) Config {
	var config Config

	viper.SetConfigFile(configFile)
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic("Failed to read config file")
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic("Failed to unmarshal config")
	}

	fmt.Println(config)

	return config
}
