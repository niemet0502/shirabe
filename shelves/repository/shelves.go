package repository

import (
	"database/sql"
)

type ShelvesRepository struct {
	db *sql.DB
}

func NewShelvesRepository(db *sql.DB) *ShelvesRepository {
	return &ShelvesRepository{db}
}
