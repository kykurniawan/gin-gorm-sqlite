package database

import (
	"log"

	"gin-gorm-sqlite/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var error error

	DB, error = gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})

	if error != nil {
		log.Fatal("failed to connect to database", error.Error())
	}

	DB.AutoMigrate(&models.Note{})
}
