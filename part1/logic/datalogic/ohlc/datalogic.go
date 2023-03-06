package fs

import (
	"context"

	"github.com/ctopher7/gltc/v2/part1/model"
)

type Datalogic interface {
	CalculateOhlc(dataset []model.Stock, stockName string) (res model.Ohlc, listOfStockName []string, err error)
	GetOhlcFromCache(ctx context.Context, stockName string) (res model.Ohlc, err error)
	StoreOhlc(ctx context.Context, data map[string]model.Ohlc) error
	PublishOhlc(ctx context.Context, data []string) error
}
