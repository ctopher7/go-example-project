package impl

import (
	ohlcDl "github.com/ctopher7/gltc/v2/part1/logic/datalogic/ohlc"
)

type datalogic struct {
}

func New() ohlcDl.Datalogic {
	return &datalogic{}
}
