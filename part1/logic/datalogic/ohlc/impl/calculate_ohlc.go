package impl

import (
	"math"
	"strconv"

	"github.com/ctopher7/gltc/v2/part1/model"
)

func (d *datalogic) CalculateOhlc(dataset []model.Stock, stockName string) (res model.Ohlc, listOfStockName []string, err error) {
	mapByStockName := make(map[string][]model.Stock)
	for _, data := range dataset {
		mapByStockName[data.StockCode] = append(mapByStockName[data.StockCode], data)
	}

	for key := range mapByStockName {
		listOfStockName = append(listOfStockName, key)
	}

	if data, ok := mapByStockName[stockName]; ok {
		for idx, stock := range data {
			if idx == 0 && stock.Quantity == "" {
				price := int64(0)
				if stock.Price != "" {
					price, err = strconv.ParseInt(stock.Price, 10, 64)
					if err != nil {
						return
					}
				}
				res.PreviousPrice = price
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
				if res.OpenPrice == 0 {
					res.OpenPrice = executionPrice
				}
				if res.HighestPrice == 0 || res.HighestPrice < executionPrice {
					res.HighestPrice = executionPrice
				}
				if res.LowestPrice == 0 || res.LowestPrice > executionPrice {
					res.LowestPrice = executionPrice
				}
				res.Volume += executedQty
				res.Value += executedQty * executionPrice

			}
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

		res.AveragePrice = int64(math.Round(float64(res.Value) / float64(res.Volume)))
		res.ClosePrice = closePrice
	}

	return
}
