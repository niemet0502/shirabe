package database

import (
	"fmt"
	"log"

	"github.com/niemet0502/shirabe/shelves/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "shelvesdb"
	port     = 5432
	user     = "marius"
	password = "root"
	dbname   = "shelve_db"
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
	db.AutoMigrate(&models.Shelf{}, &models.BookShelf{})

	return db
}
