package impl

import (
	"context"
)

func (r *repository) HSet(ctx context.Context, key string, values ...interface{}) (int64, error) {
	return r.redisConnection.HSet(ctx, key, values...).Result()
}
