package model

type GetSummaryReq struct {
	StockName string
}

type Ohlc struct {
	PreviousPrice int64 `json:"previous_price"`
	OpenPrice     int64 `json:"open_price"`
	HighestPrice  int64 `json:"highest_price"`
	LowestPrice   int64 `json:"lowest_price"`
	ClosePrice    int64 `json:"close_price"`
	AveragePrice  int64 `json:"average_price"`
	Volume        int64 `json:"volume"`
	Value         int64 `json:"value"`
}
