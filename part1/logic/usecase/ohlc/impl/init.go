package impl

import (
	fsDl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/fs"
	ohlcDl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/ohlc"
	redisRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/redis"
	"github.com/ctopher7/gltc/v2/part1/logic/usecase/ohlc"
)

type usecase struct {
	fsDatalogic     fsDl.Datalogic
	ohlcDatalogic   ohlcDl.Datalogic
	redisRepository redisRepo.Repository
}

func New(
	fsDatalogic fsDl.Datalogic,
	ohlcDatalogic ohlcDl.Datalogic,
	redisRepository redisRepo.Repository,
) ohlc.Usecase {
	return &usecase{
		fsDatalogic:     fsDatalogic,
		ohlcDatalogic:   ohlcDatalogic,
		redisRepository: redisRepository,
	}
}
