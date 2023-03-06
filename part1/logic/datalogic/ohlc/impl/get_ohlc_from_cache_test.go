package impl

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/ctopher7/gltc/v2/part1/constant"
	redisRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/redis"
	"github.com/ctopher7/gltc/v2/part1/model"
	"github.com/ctopher7/gltc/v2/part1/util"
)

func Test_GetOhlcFromCache(t *testing.T) {
	redisRepoMock := new(redisRepo.MockRepository)

	type args struct {
		stockName string
	}

	req := args{
		stockName: "BBCA",
	}

	tests := []struct {
		name    string
		mock    func()
		args    args
		want    model.Ohlc
		wantErr error
	}{
		{
			name:    "fail HGetAll",
			args:    req,
			wantErr: errors.New("err HGetAll"),
			mock: func() {
				redisRepoMock.
					On("HGetAll", context.Background(), fmt.Sprintf(constant.OhlcRedisKey, "BBCA")).
					Return(nil, errors.New("err HGetAll")).
					Once()
			},
		},
		{
			name:    "no result",
			args:    req,
			wantErr: errors.New("no data on cache"),
			mock: func() {
				redisRepoMock.
					On("HGetAll", context.Background(), fmt.Sprintf(constant.OhlcRedisKey, "BBCA")).
					Return(map[string]string{}, nil).
					Once()
			},
		},
		{
			name: "success",
			args: req,
			mock: func() {
				redisRepoMock.
					On("HGetAll", context.Background(), fmt.Sprintf(constant.OhlcRedisKey, "BBCA")).
					Return(map[string]string{
						"previous_price": "8000",
						"open_price":     "8050",
						"highest_price":  "8050",
						"lowest_price":   "8050",
						"close_price":    "8050",
						"average_price":  "8050",
						"volume":         "100",
						"value":          "805000",
					}, nil).
					Once()
			},
			want: model.Ohlc{
				PreviousPrice: 8000,
				OpenPrice:     8050,
				HighestPrice:  8050,
				LowestPrice:   8050,
				ClosePrice:    8050,
				AveragePrice:  8050,
				Volume:        100,
				Value:         805000,
			},
		},
	}

	for _, tt := range tests {
		d := datalogic{
			redisRepository: redisRepoMock,
		}

		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			got, err := d.GetOhlcFromCache(context.Background(), tt.args.stockName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOhlcFromCache test failed. want: %+v, got: %+v", tt.want, got)
			}
			if !util.SameErrorMessage(err, tt.wantErr) {
				t.Errorf("GetOhlcFromCache test failed. wantErr: %+v, gotErr: %+v", tt.wantErr, err)
			}
		})
	}
}
