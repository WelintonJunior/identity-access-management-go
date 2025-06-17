package repository

import (
	"context"
	"time"

	infraestructure "github.com/WelintonJunior/identity-access-management-go/infraestructure/redis"
	"github.com/redis/go-redis/v9"
)

type AuthTokenRepository interface {
	SetRefreshToken(ctx context.Context, key string, token string, expiration time.Duration) error
	GetRefreshToken(ctx context.Context, key string) (string, error)
	DeleteRefreshToken(ctx context.Context, key string) error
}

type authTokenRepository struct {
	redisClient *redis.Client
}

func NewAuthTokenRepository() AuthTokenRepository {
	return &authTokenRepository{redisClient: infraestructure.RedisDb}
}

func (r *authTokenRepository) SetRefreshToken(ctx context.Context, key string, token string, expiration time.Duration) error {
	return r.redisClient.Set(ctx, key, token, expiration).Err()
}

func (r *authTokenRepository) GetRefreshToken(ctx context.Context, key string) (string, error) {
	return r.redisClient.Get(ctx, key).Result()
}

func (r *authTokenRepository) DeleteRefreshToken(ctx context.Context, key string) error {
	return r.redisClient.Del(ctx, key).Err()
}
