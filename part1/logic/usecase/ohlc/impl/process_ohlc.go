package impl

import (
	"context"
	"fmt"

	"github.com/ctopher7/gltc/v2/part1/model"
)

func (u *usecase) ProcessOhlc(ctx context.Context, req model.ProcessOhlc) (err error) {
	jsonSlice, err := u.fsDatalogic.GetAllNDJsonFileInDir("subsetdata")
	if err != nil {
		fmt.Println("ERROR PROCESS OHLC1, ", err)
		return
	}
	calculatedOhlc, _, err := u.ohlcDatalogic.CalculateOhlc(jsonSlice, req.StockName)
	if err != nil {
		fmt.Println("ERROR PROCESS OHLC2, ", err)
		return
	}

	stockName := req.StockName
	err = u.ohlcDatalogic.StoreOhlc(ctx, map[string]model.Ohlc{
		stockName: calculatedOhlc,
	})
	if err != nil {
		fmt.Println("ERROR PROCESS OHLC3, ", err)
	}
	return
}
