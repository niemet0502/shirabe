package models

import (
	"gorm.io/gorm"
)

type Shelf struct {
	gorm.Model
	ID     uint `gorm:"primaryKey"`
	Name   string
	UserId int32
}
