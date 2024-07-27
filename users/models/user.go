package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	UserName string
	Email    string `gorm:"unique"`
	Password string
}
