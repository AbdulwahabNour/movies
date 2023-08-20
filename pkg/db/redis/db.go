package redis

import (
	"context"
	"time"

	"github.com/AbdulwahabNour/movies/config"
	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	Client *redis.Client
}

func NewRedisDB(config *config.Config) *RedisDB {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	redisHost := config.Redis.RedisAddr

	if redisHost == "" {
		redisHost = ":6379"
	}
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	return &RedisDB{
		Client: client,
	}
}
