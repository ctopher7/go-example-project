package part1

import (
	fsRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/fs"
	fsRepoImpl "github.com/ctopher7/gltc/v2/part1/logic/repository/fs/impl"

	fsDl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/fs"
	fsDlImpl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/fs/impl"

	ohlcUc "github.com/ctopher7/gltc/v2/part1/logic/usecase/ohlc"
	ohlcUcImpl "github.com/ctopher7/gltc/v2/part1/logic/usecase/ohlc/impl"
)

var (
	fsRepository fsRepo.Repository

	fsDatalogic fsDl.Datalogic

	ohlcUsecase ohlcUc.Usecase
)

func initRepositoryLayer() {
	fsRepository = fsRepoImpl.New()
}

func initDatalogicLayer() {
	fsDatalogic = fsDlImpl.New(fsRepository)
}

func initUsecaseLayer() {
	ohlcUsecase = ohlcUcImpl.New(fsDatalogic)
}

func initLogic() {
	initRepositoryLayer()
	initDatalogicLayer()
	initUsecaseLayer()
}
