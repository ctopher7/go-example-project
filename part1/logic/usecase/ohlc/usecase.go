package ohlc

import (
	"github.com/ctopher7/gltc/v2/part1/model"
)

type Usecase interface {
	GetSummary(req model.GetSummaryReq) (res model.GetSummaryRes, err error)
}
