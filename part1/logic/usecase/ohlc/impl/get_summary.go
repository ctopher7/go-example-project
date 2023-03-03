package impl

import "fmt"

func (u *usecase) PopulateData() error {
	jsonSlice, err := u.fsDatalogic.GetAllNDJsonFileInDir("subsetdata")
	if err != nil {
		return err
	}
	for _, json := range jsonSlice {
		fmt.Println(json)
	}
	return nil
}
