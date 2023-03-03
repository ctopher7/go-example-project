package impl

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func (r *repository) ReadNDJson(path string) (res []map[string]interface{}, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}

	d := json.NewDecoder(strings.NewReader(string(byteValue)))
	for {
		var json map[string]interface{}
		err = d.Decode(&json)
		if err != nil {
			//io.EOF is the expected err value for end of file
			if err != io.EOF {
				return
			}
			err = nil
			break
		}
		res = append(res, json)
	}

	return
}
