package handler

import (
	"context"

	ohlcUc "github.com/ctopher7/gltc/v2/part1/logic/usecase/ohlc"
	"github.com/ctopher7/gltc/v2/part1/model"
	pb "github.com/ctopher7/gltc/v2/part1/proto"
)

type GrpcHandler struct {
	pb.UnimplementedOhlcServer
	OhlcUsecase ohlcUc.Usecase
}

func (h *GrpcHandler) GetSummary(ctx context.Context, in *pb.GetSummaryRequest) (out *pb.GetSummaryResponse, err error) {
	out = &pb.GetSummaryResponse{}

	res, err := h.OhlcUsecase.GetSummary(ctx, model.GetSummaryReq{
		StockName: in.GetStockName(),
	})
	if err != nil {
		return
	}

	out.PreviousPrice = res.PreviousPrice
	out.OpenPrice = res.OpenPrice
	out.HighestPrice = res.HighestPrice
	out.LowestPrice = res.LowestPrice
	out.ClosePrice = res.ClosePrice
	out.AveragePrice = res.AveragePrice
	out.Volume = res.Volume
	out.Value = res.Value

	return
}
