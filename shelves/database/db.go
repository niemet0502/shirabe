package database

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 9438
	user     = "marius"
	password = "root"
	dbname   = "shelve_db"
)

func InitDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Print("DB connection has failed")
	}

	return db
}
