package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	UserName string `gorm:"unique"`
	Email    string
	Password string
}
