package impl

import (
	"strconv"

	"github.com/ctopher7/gltc/v2/part1/model"
)

func (d *datalogic) CalculateOhlc(dataset []model.Stock) (res map[string]model.Ohlc, err error) {
	mapByStockName := make(map[string][]model.Stock)
	res = make(map[string]model.Ohlc)
	for _, data := range dataset {
		mapByStockName[data.StockCode] = append(mapByStockName[data.StockCode], data)
	}

	for stockName, data := range mapByStockName {
		for _, stock := range data {
			executionPrice := int64(0)
			if stock.ExecutionPrice != "" {
				executionPrice, err = strconv.ParseInt(stock.ExecutionPrice, 10, 64)
				if err != nil {
					return
				}
			}
			ohlc := res[stockName]
			ohlc.OpenPrice = executionPrice

			res[stockName] = ohlc
		}
	}

	return
}
