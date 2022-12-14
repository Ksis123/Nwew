package entity

import (
	"gorm.io/gorm"
)


type User struct {
	gorm.Model

	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string `json:"-"`
	Role	 string


	Students []Student `grom:"foreignKey:UserId"`
}