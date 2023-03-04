package fs

import "github.com/ctopher7/gltc/v2/part1/model"

type Datalogic interface {
	GetAllNDJsonFileInDir(directoryPath string) (res []model.Stock, err error)
}
