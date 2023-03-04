package model

type GetSummaryReq struct {
	StockName string
}

type Ohlc struct {
	PreviousPrice int64
	OpenPrice     int64
	HighestPrice  int64
	LowestPrice   int64
	ClosePrice    int64
	AveragePrice  int64
	Volume        int64
	Value         int64
}
