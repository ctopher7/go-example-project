package impl

import (
	io "github.com/ctopher7/gltc/v2/part1/logic/repository/io"
)

type repository struct {
}

func New() io.Repository {
	return &repository{}
}
