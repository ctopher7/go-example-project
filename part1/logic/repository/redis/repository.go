package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Repository interface {
	SetNx(ctx context.Context, key string, value interface{}, expiry int) *redis.BoolCmd
}
