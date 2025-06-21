package config

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Config struct {
	DB  *gorm.DB
	RDB *redis.Client
}

func Init() (*Config, error) {
	db, err := InitDB()
	if err != nil {
		return nil, err
	}

	rdb, err := InitRedis()
	if err != nil {
		return nil, err
	}

	return &Config{
		DB:  db,
		RDB: rdb,
	}, nil

}
