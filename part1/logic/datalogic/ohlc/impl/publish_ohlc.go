package impl

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ctopher7/gltc/v2/part1/constant"
	"github.com/ctopher7/gltc/v2/part1/model"
)

func (d *datalogic) PublishOhlc(ctx context.Context, data []string) error {
	for _, stockName := range data {
		json, err := json.Marshal(model.ProcessOhlc{
			StockName: stockName,
		})
		if err != nil {
			return err
		}
		_, _, err = d.mqRepository.SendMessage(ctx, constant.ProcessOhlcTopic, string(json))
		if err != nil {
			fmt.Println("============,,", err)

			return err
		}
	}

	return nil
}
