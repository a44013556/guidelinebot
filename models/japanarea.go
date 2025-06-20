package models

import "gorm.io/gorm"

type Japanarea struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	BaseModel
}

func CheckJapanAreaExists(db *gorm.DB, name string) (bool, error) {
	var count int64
	if err := db.Model(&Japanarea{}).Where("name = ?", name).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetAllJapanAreaName(db *gorm.DB) ([]string, error) {
	var names []string
	if err := db.Model(&Japanarea{}).Pluck("name", &names).Error; err != nil {
		return nil, err
	}
	return names, nil
}
