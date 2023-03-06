package impl

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/ctopher7/gltc/v2/part1/constant"
	"github.com/ctopher7/gltc/v2/part1/model"
)

func (d *datalogic) GetOhlcFromCache(ctx context.Context, stockName string) (res model.Ohlc, err error) {
	cache, err := d.redisRepository.HGetAll(ctx, fmt.Sprintf(constant.OhlcRedisKey, stockName))
	if err != nil {
		return
	}
	if len(cache) == 0 {
		err = errors.New("no data on cache")
		return
	}
	var (
		PreviousPrice int64
		OpenPrice     int64
		HighestPrice  int64
		LowestPrice   int64
		ClosePrice    int64
		AveragePrice  int64
		Volume        int64
		Value         int64
	)

	if val, ok := cache["previous_price"]; ok && val != "" {
		PreviousPrice, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
	}
	if val, ok := cache["open_price"]; ok && val != "" {
		OpenPrice, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
	}
	if val, ok := cache["highest_price"]; ok && val != "" {
		HighestPrice, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
	}
	if val, ok := cache["lowest_price"]; ok && val != "" {
		LowestPrice, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
	}
	if val, ok := cache["close_price"]; ok && val != "" {
		ClosePrice, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
	}
	if val, ok := cache["average_price"]; ok && val != "" {
		AveragePrice, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
	}
	if val, ok := cache["volume"]; ok && val != "" {
		Volume, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
	}
	if val, ok := cache["value"]; ok && val != "" {
		Value, err = strconv.ParseInt(val, 10, 64)
		if err != nil {
			return
		}
	}
	if err != nil {
		return
	}

	res = model.Ohlc{
		PreviousPrice: PreviousPrice,
		OpenPrice:     OpenPrice,
		HighestPrice:  HighestPrice,
		LowestPrice:   LowestPrice,
		ClosePrice:    ClosePrice,
		AveragePrice:  AveragePrice,
		Volume:        Volume,
		Value:         Value,
	}
	return
}
