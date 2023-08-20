package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/AbdulwahabNour/movies/internal/token"
	"github.com/redis/go-redis/v9"
)

type redisToken struct {
	Redis *redis.Client
}

func NewTokenRepo(redisClient *redis.Client) token.TokenRepository {
	return &redisToken{
		Redis: redisClient,
	}
}

func (r *redisToken) SetToken(ctx context.Context, userID int64, token string, prefix string, expiresIn time.Duration) error {
	key := r.refreshTokenKey(prefix, userID)
	err := r.Redis.SetNX(ctx, key, token, expiresIn).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *redisToken) UpdateToken(ctx context.Context, userID int64, token string, prefix string, expiresIn time.Duration) error {
	key := r.refreshTokenKey(prefix, userID)
	err := r.Redis.SetXX(ctx, key, token, expiresIn).Err() //if the key already exists

	if err != nil {
		return err
	}

	return nil
}
func (r *redisToken) GetToken(ctx context.Context, userID int64, prefix string) (string, error) {
	key := r.refreshTokenKey(prefix, userID)
	out, err := r.Redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("invalid Token")
	}
	if err != nil {

		return "", err
	}
	return out, nil
}
func (r *redisToken) DeleteToken(ctx context.Context, userID int64, prefix string) error {
	key := r.refreshTokenKey(prefix, userID)
	out := r.Redis.Del(ctx, key)
	if err := out.Err(); err != nil {
		return err
	}
	return nil
}

func (r *redisToken) refreshTokenKey(prefix string, id int64) string {
	return fmt.Sprintf("%s:%d", prefix, id)
}
