package impl

import (
	fsDl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/fs"
	"github.com/ctopher7/gltc/v2/part1/logic/usecase/ohlc"
)

type usecase struct {
	fsDatalogic fsDl.Datalogic
}

func New(
	fsDatalogic fsDl.Datalogic,
) ohlc.Usecase {
	return &usecase{
		fsDatalogic: fsDatalogic,
	}
}
