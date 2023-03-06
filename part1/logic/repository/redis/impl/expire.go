package impl

import (
	"context"
	"time"
)

func (r *repository) Expire(ctx context.Context, key string, expiry time.Duration) (bool, error) {
	return r.redisConnection.Expire(ctx, key, expiry).Result()
}
