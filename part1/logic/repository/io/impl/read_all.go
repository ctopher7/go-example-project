package impl

import (
	"io/ioutil"
	"os"
)

func (r *repository) ReadAll(file *os.File) ([]byte, error) {
	return ioutil.ReadAll(file)
}
