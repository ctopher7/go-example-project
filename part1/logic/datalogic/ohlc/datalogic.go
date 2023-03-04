package fs

import "github.com/ctopher7/gltc/v2/part1/model"

type Datalogic interface {
	CalculateOhlc(dataset []model.Stock) (res map[string]model.Ohlc, err error)
}
