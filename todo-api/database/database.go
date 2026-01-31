package database

import (
	"log"
	"todo-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB")
	}

	DB.AutoMigrate(&models.User{}, &models.Todo{})
}
