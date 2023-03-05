package impl

import (
	"os"
)

func (r *repository) OsClose(file *os.File) error {
	return file.Close()
}
