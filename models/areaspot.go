package models

import "gorm.io/gorm"

type AreaSpot struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	AreaId       uint
	Area         JapanArea `gorm:"foreignKey:AreaId"`
	VisitedTimes uint
	Rating       float64
	BaseModel
}

func GetAreaSpotListByAreaId(db *gorm.DB, areaId int64) ([]AreaSpot, error) {
	var spots []AreaSpot
	if err := db.Model(&AreaSpot{}).Where("area_id = ?", areaId).Limit(10).Find(&spots).Error; err != nil {
		return nil, err
	}
	return spots, nil
}
