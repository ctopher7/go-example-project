package impl

import (
	"context"

	"github.com/ctopher7/gltc/v2/part1/model"
)

func (u *usecase) GetSummary(ctx context.Context, req model.GetSummaryReq) (res model.Ohlc, err error) {
	//get from cache
	//if cache error, just get directly from source of truth
	cache, err := u.ohlcDatalogic.GetOhlcFromCache(ctx, req.StockName)
	if err == nil {
		res = cache
		return
	}

	jsonSlice, err := u.fsDatalogic.GetAllNDJsonFileInDir("subsetdata")
	if err != nil {
		return
	}
	ohlc, listOfStockName, err := u.ohlcDatalogic.CalculateOhlc(jsonSlice, req.StockName)
	if err != nil {
		return
	}

	res = ohlc

	u.ohlcDatalogic.PublishOhlc(ctx, listOfStockName)
	return
}
