package impl

import (
	"errors"
	"reflect"
	"testing"

	"github.com/ctopher7/gltc/v2/part1/model"
	"github.com/ctopher7/gltc/v2/part1/util"
)

func Test_CalculateOhlc(t *testing.T) {
	type args struct {
		dataset   []model.Stock
		stockName string
	}

	tests := []struct {
		name    string
		mock    func()
		args    args
		want    model.Ohlc
		want1   []string
		wantErr error
	}{
		{
			name: "success",
			args: args{
				dataset: []model.Stock{
					{
						TransactionType:  "A",
						OrderBook:        "102",
						Price:            "8000",
						StockCode:        "BBCA",
						Quantity:         "",
						ExecutedQuantity: "",
						ExecutionPrice:   "",
						OrderVerb:        "",
					},
					{
						TransactionType:  "E",
						OrderBook:        "102",
						Price:            "",
						StockCode:        "BBCA",
						Quantity:         "",
						ExecutedQuantity: "100",
						ExecutionPrice:   "8050",
						OrderVerb:        "",
					},
				},
				stockName: "BBCA",
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
			want1: []string{"BBCA"},
		},
		{
			name: "err execution price",
			args: args{
				dataset: []model.Stock{
					{
						TransactionType:  "A",
						OrderBook:        "102",
						Price:            "8000",
						StockCode:        "BBCA",
						Quantity:         "",
						ExecutedQuantity: "",
						ExecutionPrice:   "",
						OrderVerb:        "",
					},
					{
						TransactionType:  "E",
						OrderBook:        "102",
						Price:            "",
						StockCode:        "BBCA",
						Quantity:         "",
						ExecutedQuantity: "100",
						ExecutionPrice:   "asd",
						OrderVerb:        "",
					},
				},
				stockName: "BBCA",
			},
			wantErr: errors.New(`strconv.ParseInt: parsing "asd": invalid syntax`),
			want1:   []string{"BBCA"},
			want: model.Ohlc{
				PreviousPrice: 8000,
			},
		},
		{
			name: "err executed qty",
			args: args{
				dataset: []model.Stock{
					{
						TransactionType:  "A",
						OrderBook:        "102",
						Price:            "8000",
						StockCode:        "BBCA",
						Quantity:         "",
						ExecutedQuantity: "",
						ExecutionPrice:   "",
						OrderVerb:        "",
					},
					{
						TransactionType:  "E",
						OrderBook:        "102",
						Price:            "",
						StockCode:        "BBCA",
						Quantity:         "",
						ExecutedQuantity: "asd",
						ExecutionPrice:   "8050",
						OrderVerb:        "",
					},
				},
				stockName: "BBCA",
			},
			wantErr: errors.New(`strconv.ParseInt: parsing "asd": invalid syntax`),
			want1:   []string{"BBCA"},
			want: model.Ohlc{
				PreviousPrice: 8000,
			},
		},
	}

	for _, tt := range tests {
		d := datalogic{}

		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			got, got1, err := d.CalculateOhlc(tt.args.dataset, tt.args.stockName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateOhlc test failed. want: %+v, got: %+v", tt.want, got)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CalculateOhlc test failed. want1: %+v, got1: %+v", tt.want1, got1)
			}
			if !util.SameErrorMessage(err, tt.wantErr) {
				t.Errorf("CalculateOhlc test failed. wantErr: %+v, gotErr: %+v", tt.wantErr, err)
			}
		})
	}
}
