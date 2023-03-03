package fs

type Datalogic interface {
	GetAllNDJsonFileInDir(directoryPath string) (res []map[string]interface{}, err error)
}
