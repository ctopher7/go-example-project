package impl

import (
	fsDl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/fs"
	ioRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/io"
	jsonRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/json"
)

type datalogic struct {
	io   ioRepo.Repository
	json jsonRepo.Repository
}

func New(
	io ioRepo.Repository,
	json jsonRepo.Repository,
) fsDl.Datalogic {
	return &datalogic{
		io:   io,
		json: json,
	}
}
