package ohlc

import (
	"context"

	"github.com/ctopher7/gltc/v2/part1/model"
)

type Usecase interface {
	GetSummary(req model.GetSummaryReq) (res model.Ohlc, err error)
	CalculateData(ctx context.Context) error
}
