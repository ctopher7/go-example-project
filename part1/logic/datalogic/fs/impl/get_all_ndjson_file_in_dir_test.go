package impl

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"time"

	ioRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/io"
	jsonRepo "github.com/ctopher7/gltc/v2/part1/logic/repository/json"
	"github.com/ctopher7/gltc/v2/part1/model"
	"github.com/ctopher7/gltc/v2/part1/util"
)

type FileInfo struct {
	FileName string
}

func (f FileInfo) Name() string {
	return f.FileName
}

func (f FileInfo) Size() (res int64)        { return }
func (f FileInfo) Mode() (res fs.FileMode)  { return }
func (f FileInfo) ModTime() (res time.Time) { return }
func (f FileInfo) IsDir() (res bool)        { return }
func (f FileInfo) Sys() (res any)           { return }

func Test_GetAllNDJsonFileInDir(t *testing.T) {
	ioRepoMock := new(ioRepo.MockRepository)
	jsonRepoMock := new(jsonRepo.MockRepository)

	type args struct {
		directoryPath string
	}

	req := args{
		directoryPath: "test",
	}

	tests := []struct {
		name    string
		mock    func()
		args    args
		want    []model.Stock
		wantErr error
	}{
		{
			name: "fail ReadDir",
			args: req,
			mock: func() {
				ioRepoMock.
					On("ReadDir", "test").
					Return(nil, errors.New("err read dir")).
					Once()
			},
			wantErr: errors.New("err read dir"),
		},
		{
			name: "fail OsOpen",
			args: req,
			mock: func() {
				ioRepoMock.
					On("ReadDir", "test").
					Return([]fs.FileInfo{FileInfo{FileName: "name"}}, nil).
					Once()

				ioRepoMock.
					On("OsOpen", "test/name").
					Return(nil, errors.New("err os open")).
					Once()
			},
			wantErr: errors.New("err os open"),
		},
	}

	for _, tt := range tests {
		d := datalogic{
			io:   ioRepoMock,
			json: jsonRepoMock,
		}

		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			got, err := d.GetAllNDJsonFileInDir(tt.args.directoryPath)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllNDJsonFileInDir test failed. want: %+v, got: %+v", tt.want, got)
			}
			if !util.SameErrorMessage(err, tt.wantErr) {
				t.Errorf("GetAllNDJsonFileInDir test failed. wantErr: %+v, gotErr: %+v", tt.wantErr, err)
			}
		})
	}
}
