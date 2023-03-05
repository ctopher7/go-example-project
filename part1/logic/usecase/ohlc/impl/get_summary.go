package impl

import (
	"context"
	"fmt"

	"github.com/ctopher7/gltc/v2/part1/model"
)

func (u *usecase) GetSummary(ctx context.Context, req model.GetSummaryReq) (res model.Ohlc, err error) {
	jsonSlice, err := u.fsDatalogic.GetAllNDJsonFileInDir("subsetdata")
	if err != nil {
		return
	}
	ohlc, err := u.ohlcDatalogic.CalculateOhlc(jsonSlice)
	if err != nil {
		return
	}
	fmt.Printf("%+v \n", ohlc)
	u.CalculateData(ctx)
	return
}
