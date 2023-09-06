package config

import (
	"fmt"
	"golang-rest-api/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	config Config
}

func NewDB(cfg Config) *DB {
	return &DB{config: cfg}
}

func (db *DB) ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", db.config.Database.Host, db.config.Database.Username, db.config.Database.Password, db.config.Database.DBName, db.config.Database.Port)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if err := conn.AutoMigrate(&domain.Employee{}, &domain.Customer{}, &domain.Product{}, &domain.Order{}, &domain.OrderDetail{}, &domain.Transaction{}); err != nil {
		panic(err)
	}

	return conn
}
