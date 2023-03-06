package part1

import (
	"github.com/Shopify/sarama"
	"github.com/redis/go-redis/v9"
)

type Resource struct {
	RedisConnection *redis.Client
	MqProducer      sarama.SyncProducer
}

func InitResource() Resource {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "testing123",
		DB:       0,
	})

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true
	config.Producer.MaxMessageBytes = 200000000

	producer, err := sarama.NewSyncProducer([]string{"localhost:19092"}, config)
	if err != nil {
		panic(err)
	}

	return Resource{
		RedisConnection: rdb,
		MqProducer:      producer,
	}
}
