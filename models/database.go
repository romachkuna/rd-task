package models

import (
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
