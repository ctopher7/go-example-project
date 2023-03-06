package ohlc

import (
	"context"

	"github.com/ctopher7/gltc/v2/part1/model"
)

type Usecase interface {
	GetSummary(ctx context.Context, req model.GetSummaryReq) (res model.Ohlc, err error)
	ProcessOhlc(ctx context.Context, req model.ProcessOhlc) (err error)
}
