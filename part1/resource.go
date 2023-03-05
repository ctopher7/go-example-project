package part1

import "github.com/redis/go-redis/v9"

type Resource struct {
	RedisConnection *redis.Client
}

func InitResource() Resource {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "testing123",
		DB:       0,
	})

	return Resource{
		RedisConnection: rdb,
	}
}
