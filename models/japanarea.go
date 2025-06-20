package models

import "gorm.io/gorm"

type JapanArea struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	BaseModel
}

func CheckJapanAreaExists(db *gorm.DB, name string) (*Japanarea, error) {
	var area Japanarea
	if err := db.Model(&Japanarea{}).Where("name = ?", name).First(&area).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &area, nil
}

func GetAllJapanAreaName(db *gorm.DB) ([]string, error) {
	var names []string
	if err := db.Model(&Japanarea{}).Pluck("name", &names).Error; err != nil {
		return nil, err
	}
	return names, nil
}
