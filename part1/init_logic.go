package part1

import (
	fsRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/fs"
	fsRepoImpl "github.com/ctopher7/gltc/v2/part1/logic/repository/fs/impl"

	fsDl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/fs"
	fsDlImpl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/fs/impl"
	ohlcDl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/ohlc"
	ohlcDlImpl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/ohlc/impl"

	ohlcUc "github.com/ctopher7/gltc/v2/part1/logic/usecase/ohlc"
	ohlcUcImpl "github.com/ctopher7/gltc/v2/part1/logic/usecase/ohlc/impl"
)

var (
	fsRepository fsRepo.Repository

	fsDatalogic   fsDl.Datalogic
	ohlcDatalogic ohlcDl.Datalogic

	ohlcUsecase ohlcUc.Usecase
)

func initRepositoryLayer() {
	fsRepository = fsRepoImpl.New()
}

func initDatalogicLayer() {
	fsDatalogic = fsDlImpl.New(fsRepository)
	ohlcDatalogic = ohlcDlImpl.New()
}

func initUsecaseLayer() {
	ohlcUsecase = ohlcUcImpl.New(fsDatalogic, ohlcDatalogic)
}

func initLogic() {
	initRepositoryLayer()
	initDatalogicLayer()
	initUsecaseLayer()
}
