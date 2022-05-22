package configs

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("bookify.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection error", err)
	}
}
