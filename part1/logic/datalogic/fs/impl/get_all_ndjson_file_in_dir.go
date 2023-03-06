package impl

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"sort"
	"strings"

	"github.com/ctopher7/gltc/v2/part1/model"
)

func (d *datalogic) GetAllNDJsonFileInDir(directoryPath string) (res []model.Stock, err error) {
	//first get all file names in subsetdata dir
	files, err := d.io.ReadDir(directoryPath)
	if err != nil {
		return
	}

	tasks := make(chan model.AsyncTaskRes)
	//read every files using goroutine
	for idx, file := range files {
		go d.readFiles(idx, file, directoryPath, tasks)
	}

	tasksRes := []model.AsyncTaskRes{}
	for range files {
		got := <-tasks
		if got.Err != nil {
			err = got.Err
			return
		}

		tasksRes = append(tasksRes, got)
	}

	//sort based on file index
	sort.Slice(tasksRes, func(i, j int) bool {
		return tasksRes[i].Idx < tasksRes[j].Idx
	})

	for _, task := range tasksRes {
		if val, ok := task.Res.([]model.Stock); ok {
			for _, stock := range val {
				res = append(res, stock)
			}
		}
	}

	return
}

func (d *datalogic) readFiles(idx int, file fs.FileInfo, directoryPath string, out chan<- model.AsyncTaskRes) {
	var (
		jsonFile  *os.File
		byteValue []byte
		err       error
		res       []model.Stock
	)
	jsonFile, err = d.io.OsOpen(fmt.Sprintf("%s/%s", directoryPath, file.Name()))
	if err != nil {
		out <- model.AsyncTaskRes{
			Err: err,
			Idx: idx,
			Res: res,
		}
	}
	defer d.io.OsClose(jsonFile)

	byteValue, err = d.io.ReadAll(jsonFile)
	if err != nil {
		out <- model.AsyncTaskRes{
			Err: err,
			Idx: idx,
			Res: res,
		}
	}

	got := json.NewDecoder(strings.NewReader(string(byteValue)))
	for {
		var json model.Stock
		err = d.json.Decode(got, &json)
		if err != nil {
			//io.EOF is the expected err value for end of file
			if err != io.EOF {
				out <- model.AsyncTaskRes{
					Err: err,
					Idx: idx,
					Res: res,
				}
			}
			err = nil
			break
		}
		res = append(res, json)
	}
	out <- model.AsyncTaskRes{
		Err: err,
		Idx: idx,
		Res: res,
	}
}
