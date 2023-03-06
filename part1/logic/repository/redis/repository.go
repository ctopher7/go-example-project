package redis

import (
	"context"
	"time"
)

type Repository interface {
	HSet(ctx context.Context, key string, values ...interface{}) (int64, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	Expire(ctx context.Context, key string, expiry time.Duration) (bool, error)
}
