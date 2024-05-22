package models

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type PostgresDB struct {
	instance *gorm.DB
}

var DB *PostgresDB

func NewPostgresDB() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	dsn := os.Getenv("dsn")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{QueryFields: true})
	if err != nil {
		log.Fatal(err)
		return err
	}
	DB = &PostgresDB{
		instance: db,
	}
	err = db.AutoMigrate(&Request{}, &Response{})
	if err != nil {
		return err
	}
	return nil
}
