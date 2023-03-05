package impl

import (
	ohlcDl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/ohlc"
	redisRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/redis"
)

type datalogic struct {
	redisRepository redisRepo.Repository
}

func New(
	redisRepository redisRepo.Repository,
) ohlcDl.Datalogic {
	return &datalogic{
		redisRepository: redisRepository,
	}
}
