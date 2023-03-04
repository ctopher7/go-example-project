package impl

import (
	"math"
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
		for idx, stock := range data {
			ohlc := res[stockName]
			if idx == 0 && stock.Quantity == "" {
				price := int64(0)
				if stock.Price != "" {
					price, err = strconv.ParseInt(stock.Price, 10, 64)
					if err != nil {
						return
					}
				}
				ohlc.PreviousPrice = price
				res[stockName] = ohlc
				continue
			}

			executionPrice := int64(0)
			if stock.ExecutionPrice != "" {
				executionPrice, err = strconv.ParseInt(stock.ExecutionPrice, 10, 64)
				if err != nil {
					return
				}
			}

			executedQty := int64(0)
			if stock.ExecutedQuantity != "" {
				executedQty, err = strconv.ParseInt(stock.ExecutedQuantity, 10, 64)
				if err != nil {
					return
				}
			}

			if stock.TransactionType == "E" || stock.TransactionType == "P" {
				if ohlc.OpenPrice == 0 {
					ohlc.OpenPrice = executionPrice
				}
				if ohlc.HighestPrice == 0 || ohlc.HighestPrice < executionPrice {
					ohlc.HighestPrice = executionPrice
				}
				if ohlc.LowestPrice == 0 || ohlc.LowestPrice > executionPrice {
					ohlc.LowestPrice = executionPrice
				}
				ohlc.Volume += executedQty
				ohlc.Value += executedQty * executionPrice

			}

			res[stockName] = ohlc
		}

		closePrice := int64(0)

		for i := len(data) - 1; i >= 0; i-- {
			if (data[i].TransactionType == "E" || data[i].TransactionType == "P") && data[i].ExecutionPrice != "" {
				closePrice, err = strconv.ParseInt(data[i].ExecutionPrice, 10, 64)
				if err != nil {
					return
				}
				break
			}
		}

		ohlc := res[stockName]
		ohlc.AveragePrice = int64(math.Round(float64(ohlc.Value) / float64(ohlc.Volume)))
		ohlc.ClosePrice = closePrice
		res[stockName] = ohlc
	}

	return
}
