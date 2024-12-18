package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID              uint `gorm:"primaryKey"`
	Title           string
	Author          string
	Slug            string
	Genre           string
	Status          int
	TotalPages      int
	ReadingProgress int
	UserId          int `gorm:"not null"`
	Cover           string
	Description     string
	PublicatedAt    time.Time `gorm:"default:null"`
	StartReadingAt  time.Time `gorm:"default:null"`
	FinishReadingAt time.Time `gorm:"default:null"`
}

type CreateBook struct {
	Title        string
	Author       string
	Genre        string
	TotalPages   int
	UserId       int
	Cover        string
	Description  string
	PublicatedAt time.Time
}
