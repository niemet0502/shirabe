package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/niemet0502/shirabe/books/models"
)

const (
	host     = "localhost"
	port     = 9439
	user     = "marius"
	password = "root"
	dbname   = "book_db"
)

func InitDb() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s user=%s  "+
		"password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.Book{})

	return db
}
