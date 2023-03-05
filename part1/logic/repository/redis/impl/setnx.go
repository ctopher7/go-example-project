package impl

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r *repository) SetNx(ctx context.Context, key string, value interface{}, expiry int) *redis.BoolCmd {
	return r.redisConnection.SetNX(ctx, key, value, time.Duration(expiry))
}
