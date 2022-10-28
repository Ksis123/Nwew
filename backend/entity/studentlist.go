package entity

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	Name         string
	StudentLists []StudentList `gorm:"foreignKey:ReportID"`
}
type Status struct {
	gorm.Model

	Status       string

	StudentLists []StudentList `gorm:"foreignKey:StatusID"`
}
type StudentList struct {
	gorm.Model

	Reason string
	Amount int
	SaveTime  time.Time

	AdminID *uint
	Admin   Admin

	ReportID *uint
	Report   Report 

	StatusID *uint
	Status   Status 
	
	
	SlipLists []SlipList `gorm:"foreignKey:StudentListID"`
}
