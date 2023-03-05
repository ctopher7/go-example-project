package impl

import (
	"io/fs"
	"io/ioutil"
)

func (r *repository) ReadDir(dirname string) ([]fs.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}
