package models

import (
	"gorm.io/gorm"
)

type ReadingProgress struct {
	gorm.Model

	ID          uint `gorm:"primaryKey"`
	UserId      int  `gorm:"not null"`
	BookId      int  `gorm:"not null"`
	PagesNumber int
}

type CreateReadingProgress struct {
	UserId      int
	BookId      int
	PagesNumber int
}
