package models

import (
	"time"
)

type Booking struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Date       string
	People     int
	Areas      []Japanarea
	CancelDate *time.Time `gorm:"column:cancel_date"`
	BaseModel
}
