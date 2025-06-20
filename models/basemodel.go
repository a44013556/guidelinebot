package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `gorm:"column:create_date;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:update_date;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_date"`
}
