package impl

import (
	redisRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/redis"
	"github.com/redis/go-redis/v9"
)

type repository struct {
	redisConnection *redis.Client
}

func New(
	redisConnection *redis.Client,
) redisRepo.Repository {
	return &repository{
		redisConnection: redisConnection,
	}
}
