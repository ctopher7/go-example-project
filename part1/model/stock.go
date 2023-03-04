package model

type Stock struct {
	TransactionType  string `json:"type"`
	OrderBook        string `json:"order_book"`
	Price            string `json:"price"`
	StockCode        string `json:"stock_code"`
	ExecutedQuantity string `json:"executed_quantity"`
	ExecutionPrice   string `json:"execution_price"`
	OrderVerb        string `json:"order_verb"`
}
