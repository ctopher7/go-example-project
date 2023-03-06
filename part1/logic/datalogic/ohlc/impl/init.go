package impl

import (
	ohlcDl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/ohlc"
	mqRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/mq"
	redisRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/redis"
)

type datalogic struct {
	redisRepository redisRepo.Repository
	mqRepository    mqRepo.Repository
}

func New(
	redisRepository redisRepo.Repository,
	mqRepository mqRepo.Repository,
) ohlcDl.Datalogic {
	return &datalogic{
		redisRepository: redisRepository,
		mqRepository:    mqRepository,
	}
}
