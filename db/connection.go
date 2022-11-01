package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DNS = "host=localhost user=juliovdv password=123456 dbname=gorm port=5432"
var DB *gorm.DB

func DBconnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	}
}
