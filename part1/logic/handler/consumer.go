package handler

import (
	"context"
	"encoding/json"

	ohlcUc "github.com/ctopher7/gltc/v2/part1/logic/usecase/ohlc"
	"github.com/ctopher7/gltc/v2/part1/model"
)

type ConsumerHandler struct {
	OhlcUsecase ohlcUc.Usecase
}

func (h *ConsumerHandler) ProcessOhlc(in []byte) {
	req := model.ProcessOhlc{}
	err := json.Unmarshal(in, &req)
	if err != nil {
		return
	}
	h.OhlcUsecase.ProcessOhlc(context.Background(), req)
}
