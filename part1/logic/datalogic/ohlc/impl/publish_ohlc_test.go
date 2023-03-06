package impl

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/ctopher7/gltc/v2/part1/constant"
	mqRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/mq"
	"github.com/ctopher7/gltc/v2/part1/model"
	"github.com/ctopher7/gltc/v2/part1/util"
)

func Test_PublishOhlc(t *testing.T) {
	mqRepoMock := new(mqRepo.MockRepository)

	type args struct {
		data []string
	}

	req := args{
		data: []string{"BBCA"},
	}

	tests := []struct {
		name    string
		mock    func()
		args    args
		wantErr error
	}{
		{
			name: "err SendMessage",
			args: req,
			mock: func() {
				j, _ := json.Marshal(model.ProcessOhlc{
					StockName: "BBCA",
				})
				mqRepoMock.
					On("SendMessage", context.Background(), constant.ProcessOhlcTopic, string(j)).
					Return(int32(0), int64(0), errors.New("err SendMessage")).
					Once()
			},
			wantErr: errors.New("err SendMessage"),
		},
		{
			name: "success",
			args: req,
			mock: func() {
				j, _ := json.Marshal(model.ProcessOhlc{
					StockName: "BBCA",
				})
				mqRepoMock.
					On("SendMessage", context.Background(), constant.ProcessOhlcTopic, string(j)).
					Return(int32(0), int64(0), nil).
					Once()
			},
		},
	}

	for _, tt := range tests {
		d := datalogic{
			mqRepository: mqRepoMock,
		}

		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			err := d.PublishOhlc(context.Background(), tt.args.data)

			if !util.SameErrorMessage(err, tt.wantErr) {
				t.Errorf("PublishOhlc test failed. wantErr: %+v, gotErr: %+v", tt.wantErr, err)
			}
		})
	}
}
