package redis

import (
	"context"
)

type Repository interface {
	SendMessage(ctx context.Context, topic, value string) (int32, int64, error)
}
