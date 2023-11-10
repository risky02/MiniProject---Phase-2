package config

import (
	"Phase2/entity"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Load file ENV")
	}

	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Database")
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Checkout{}, &entity.Payment{}, &entity.Equipment{})
	if err != nil {
		return nil, err
	}

	return db, nil
}