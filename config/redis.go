package config

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

func InitRedis() (*redis.Client, error) {
	redisUrl := os.Getenv("REDIS_URL")
	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		return nil, err
	}
	rdb := redis.NewClient(opt)
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}
