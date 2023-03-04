package fs

import "github.com/ctopher7/gltc/v2/part1/model"

type Repository interface {
	GetAllFilename(directoryPath string) (res []string, err error)
	ReadNDJson(path string) (res []model.Stock, err error)
}
