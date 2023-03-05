package impl

import (
	"os"
)

func (r *repository) OsOpen(name string) (*os.File, error) {
	return os.Open(name)
}
