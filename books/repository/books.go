package repository

import (
	"database/sql"
	"fmt"
)

// the book repository needs the db connection

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (repo *BookRepository) GetBook() {
	fmt.Print("Failed to fetch the book")
}
