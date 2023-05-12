package database

import (
	"backend/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect() {
	Instance, dbError = gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) //! creates test.db in root if not found
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to db")
	}

	log.Println("Connected to databases")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	Instance.AutoMigrate(&models.Post{})
	log.Println("Database migration Completed")
}
