package impl

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ctopher7/gltc/v2/part1/constant"
	redisRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/redis"
	"github.com/ctopher7/gltc/v2/part1/model"
	"github.com/ctopher7/gltc/v2/part1/util"
)

func Test_StoreOhlc(t *testing.T) {
	redisRepoMock := new(redisRepo.MockRepository)

	type args struct {
		data map[string]model.Ohlc
	}

	req := args{
		data: map[string]model.Ohlc{
			"BBCA": {
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

	tests := []struct {
		name    string
		mock    func()
		args    args
		wantErr error
	}{
		{
			name:    "fail HSet",
			args:    req,
			wantErr: errors.New("err HSet"),
			mock: func() {
				val := []interface{}{
					context.Background(),
					fmt.Sprintf(constant.OhlcRedisKey, "BBCA"),
					"previous_price", int64(8000),
					"open_price", int64(8050),
					"highest_price", int64(8050),
					"lowest_price", int64(8050),
					"close_price", int64(8050),
					"average_price", int64(8050),
					"volume", int64(100),
					"value", int64(805000),
				}
				redisRepoMock.
					On("HSet", val...).
					Return(int64(0), errors.New("err HSet")).
					Once()
			},
		},
		{
			name: "success",
			args: req,
			mock: func() {
				val := []interface{}{
					context.Background(),
					fmt.Sprintf(constant.OhlcRedisKey, "BBCA"),
					"previous_price", int64(8000),
					"open_price", int64(8050),
					"highest_price", int64(8050),
					"lowest_price", int64(8050),
					"close_price", int64(8050),
					"average_price", int64(8050),
					"volume", int64(100),
					"value", int64(805000),
				}
				redisRepoMock.
					On("HSet", val...).
					Return(int64(0), nil).
					Once()

				redisRepoMock.
					On("Expire", context.Background(), fmt.Sprintf(constant.OhlcRedisKey, "BBCA"), 10*time.Second).
					Return(true, nil).
					Once()
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

			err := d.StoreOhlc(context.Background(), tt.args.data)

			if !util.SameErrorMessage(err, tt.wantErr) {
				t.Errorf("StoreOhlc test failed. wantErr: %+v, gotErr: %+v", tt.wantErr, err)
			}
		})
	}
}
