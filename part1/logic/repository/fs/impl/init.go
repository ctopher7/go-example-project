package impl

import (
	"github.com/ctopher7/gltc/v2/part1/logic/repository/fs"
)

type repository struct {
}

func New() fs.Repository {
	return &repository{}
}
