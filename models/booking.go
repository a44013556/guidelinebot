package models

type Booking struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Date string
}
