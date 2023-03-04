package impl

import (
	fsDl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/fs"
	ohlcDl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/ohlc"
	"github.com/ctopher7/gltc/v2/part1/logic/usecase/ohlc"
)

type usecase struct {
	fsDatalogic   fsDl.Datalogic
	ohlcDatalogic ohlcDl.Datalogic
}

func New(
	fsDatalogic fsDl.Datalogic,
	ohlcDatalogic ohlcDl.Datalogic,
) ohlc.Usecase {
	return &usecase{
		fsDatalogic:   fsDatalogic,
		ohlcDatalogic: ohlcDatalogic,
	}
}
