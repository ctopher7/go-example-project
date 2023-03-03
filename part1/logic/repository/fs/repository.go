package fs

type Repository interface {
	GetAllFilename(directoryPath string) (res []string, err error)
	ReadNDJson(path string) (res []map[string]interface{}, err error)
}
