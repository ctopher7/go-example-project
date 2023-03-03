package impl

import (
	"fmt"
	"io/ioutil"
)

func (r *repository) GetAllFilename(directoryPath string) (res []string, err error) {
	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		return
	}

	for _, file := range files {
		res = append(res, fmt.Sprintf("%s/%s", directoryPath, file.Name()))
	}
	return
}
