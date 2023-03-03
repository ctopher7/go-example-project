package impl

func (d *datalogic) GetAllNDJsonFileInDir(directoryPath string) (res []map[string]interface{}, err error) {
	files, err := d.fsRepository.GetAllFilename(directoryPath)
	if err != nil {
		return
	}
	for _, file := range files {
		json, errJson := d.fsRepository.ReadNDJson(file)
		if errJson != nil {
			err = errJson
			return
		}
		res = append(res, json...)
	}
	return
}
