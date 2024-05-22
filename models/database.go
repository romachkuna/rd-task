package models

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB PostgresDB

type PostgresDB struct {
	Instance *gorm.DB
}

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
	DB = PostgresDB{
		Instance: db,
	}
	err = db.AutoMigrate(&Request{}, &Response{})
	if err != nil {
		return err
	}
	return nil
}
