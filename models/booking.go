package models

import (
	"time"
)

type Booking struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Date       string
	People     int
	Areas      []Japanarea `gorm:"many2many:booking_areas;"`
	CancelDate *time.Time `gorm:"column:cancel_date"`
	BaseModel
}
