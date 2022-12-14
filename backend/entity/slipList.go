package entity

import (
	"time"

	"gorm.io/gorm"
)

type Banking struct {
	gorm.Model
	Name      string
	Commerce  string
	Branch    string
	SlipLists []SlipList `gorm:"foreignKey:BankingID"`
}
type PaymentStatus struct {
	gorm.Model
	Name      string
	SlipLists []SlipList `gorm:"foreignKey:PayID"`
}



type SlipList struct {
	gorm.Model
	Total    float64
	Slipdate time.Time

	AdminID *uint
	Admin   Admin

	BankingID *uint
	Banking   Banking

	PayID *uint
	Pay   PaymentStatus

	StudentListID *uint
	StudentList   StudentList
}
