package impl

import (
	"fmt"

	"github.com/ctopher7/gltc/v2/part1/model"
)

func (u *usecase) GetSummary(req model.GetSummaryReq) (res model.Ohlc, err error) {
	jsonSlice, err := u.fsDatalogic.GetAllNDJsonFileInDir("subsetdata")
	if err != nil {
		return
	}
	ohlc, err := u.ohlcDatalogic.CalculateOhlc(jsonSlice)
	if err != nil {
		return
	}
	fmt.Printf("%+v \n", ohlc)
	return
}
