package part1

import (
	ioRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/io"
	ioRepoImpl "github.com/ctopher7/gltc/v2/part1/logic/repository/io/impl"
	jsonRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/json"
	jsonRepoImpl "github.com/ctopher7/gltc/v2/part1/logic/repository/json/impl"
	mqRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/mq"
	mqRepoImpl "github.com/ctopher7/gltc/v2/part1/logic/repository/mq/impl"
	redisRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/redis"
	redisRepoImpl "github.com/ctopher7/gltc/v2/part1/logic/repository/redis/impl"

	fsDl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/fs"
	fsDlImpl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/fs/impl"
	ohlcDl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/ohlc"
	ohlcDlImpl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/ohlc/impl"

	ohlcUc "github.com/ctopher7/gltc/v2/part1/logic/usecase/ohlc"
	ohlcUcImpl "github.com/ctopher7/gltc/v2/part1/logic/usecase/ohlc/impl"
)

var (
	ioRepository    ioRepo.Repository
	jsonRepository  jsonRepo.Repository
	redisRepository redisRepo.Repository
	mqRepository    mqRepo.Repository

	fsDatalogic   fsDl.Datalogic
	ohlcDatalogic ohlcDl.Datalogic

	ohlcUsecase ohlcUc.Usecase
)

func initRepositoryLayer(resource Resource) {
	ioRepository = ioRepoImpl.New()
	jsonRepository = jsonRepoImpl.New()
	redisRepository = redisRepoImpl.New(resource.RedisConnection)
	mqRepository = mqRepoImpl.New(resource.MqProducer)
}

func initDatalogicLayer() {
	fsDatalogic = fsDlImpl.New(ioRepository, jsonRepository)
	ohlcDatalogic = ohlcDlImpl.New(redisRepository, mqRepository)
}

func initUsecaseLayer() {
	ohlcUsecase = ohlcUcImpl.New(fsDatalogic, ohlcDatalogic, redisRepository)
}

func initLogic(resource Resource) {
	initRepositoryLayer(resource)
	initDatalogicLayer()
	initUsecaseLayer()
}
