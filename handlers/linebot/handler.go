package linebot

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type LineBotHandler struct {
	DB  *gorm.DB
	RDB *redis.Client
}

func NewHandler(db *gorm.DB, rdb *redis.Client) *LineBotHandler {
	return &LineBotHandler{
		DB:  db,
		RDB: rdb,
	}
}
