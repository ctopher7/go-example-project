package impl

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/ctopher7/gltc/v2/part1/constant"
	"github.com/ctopher7/gltc/v2/part1/model"
)

func (d *datalogic) StoreOhlc(ctx context.Context, data map[string]model.Ohlc) error {
	for key, val := range data {
		valueToRedis := []interface{}{}
		structKey := reflect.ValueOf(val)
		for i := 0; i < structKey.Type().NumField(); i++ {
			valueToRedis = append(valueToRedis, structKey.Type().Field(i).Tag.Get("json"), int64(reflect.Indirect(structKey).FieldByName(structKey.Type().Field(i).Name).Int()))
		}

		redisKey := fmt.Sprintf(constant.OhlcRedisKey, key)
		_, err := d.redisRepository.HSet(ctx, redisKey, valueToRedis...)
		if err != nil {
			return err
		}

		d.redisRepository.Expire(ctx, redisKey, 10*time.Second)
	}
	return nil
}
