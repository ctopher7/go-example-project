package impl

import (
	"context"
)

func (r *repository) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.redisConnection.HGetAll(ctx, key).Result()
}
