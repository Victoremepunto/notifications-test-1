package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("notifications.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
}
