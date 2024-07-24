package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title           string
	Author          string
	Genre           string
	Status          int
	TotalPages      int
	ReadingProgress int
	UserId          int
}

type CreateBook struct {
	Title      string
	Author     string
	Genre      string
	TotalPages int
	UserId     int
}
