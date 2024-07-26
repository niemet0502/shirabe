package models

import "gorm.io/gorm"

type BookShelf struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	BookId  int32
	ShelfId int32
}
