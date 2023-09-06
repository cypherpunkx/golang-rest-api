package config

import (
	"github.com/spf13/viper"
)

// Config struktur konfigurasi aplikasi
type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
	} `mapstructure:"database"`
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

	return config
}
