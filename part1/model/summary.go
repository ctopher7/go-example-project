package model

type GetSummaryReq struct {
	StockName string
}

type GetSummaryRes struct {
	PreviousPrice int32
	OpenPrice     int32
	HighestPrice  int32
	LowestPrice   int32
	ClosePrice    int32
	AveragePrice  int32
	Volume        int32
	Value         int32
}
