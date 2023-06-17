package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {

	connStr := "host=localhost user=postgres password=root dbname=db_laundry port=5432 sslmode=disable "

	database, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}

	DB = database

}
