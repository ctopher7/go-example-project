package impl

import (
	"context"

	"github.com/ctopher7/gltc/v2/part1/model"
)

func (u *usecase) ProcessOhlc(ctx context.Context, req model.ProcessOhlc) (err error) {
	jsonSlice, err := u.fsDatalogic.GetAllNDJsonFileInDir("subsetdata")
	if err != nil {
		return
	}
	calculatedOhlc, _, err := u.ohlcDatalogic.CalculateOhlc(jsonSlice, req.StockName)
	if err != nil {
		return
	}

	stockName := req.StockName
	err = u.ohlcDatalogic.StoreOhlc(ctx, map[string]model.Ohlc{
		stockName: calculatedOhlc,
	})
	return
}
