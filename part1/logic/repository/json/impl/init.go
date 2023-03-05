package impl

import (
	"github.com/ctopher7/gltc/v2/part1/logic/repository/json"
)

type repository struct {
}

func New() json.Repository {
	return &repository{}
}
