package impl

import (
	fsDl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/fs"
	fsRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/fs"
)

type datalogic struct {
	fsRepository fsRepo.Repository
}

func New(
	fsRepository fsRepo.Repository,
) fsDl.Datalogic {
	return &datalogic{
		fsRepository: fsRepository,
	}
}
