package models

import (
	"gorm.io/gorm"
)

type Shelf struct {
	gorm.Model
	Name   string
	UserId int32
}
