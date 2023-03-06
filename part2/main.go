package part2

import (
	"encoding/json"
	"fmt"
)

type ChangeRecord struct {
	StockCode string
	Price     int64
	Quantity  int64
}

type IndexMember struct {
	StockCode string
	IndexCode string
}

type Summary struct {
	StockCode string   `json:"stock_code"`
	IndexCode []string `json:"index_code"`
	Open      int64    `json:"open"`
	High      int64    `json:"high"`
	Low       int64    `json:"low"`
	Close     int64    `json:"close"`
	Prev      int64    `json:"prev"`
}

func ohlc(records []ChangeRecord, indexMembers []IndexMember) (result map[string]Summary) {
	result = make(map[string]Summary)

	for _, record := range records {
		temp := Summary{}
		if val, ok := result[record.StockCode]; ok {
			temp = val
		}

		if record.Quantity == 0 {
			temp.Prev = record.Price
			result[record.StockCode] = temp
			continue
		}
		if temp.Open == 0 && record.Quantity > 0 {
			temp.Open = record.Price
			result[record.StockCode] = temp
			continue
		}
		temp.Close = record.Price
		if temp.High < record.Price {
			temp.High = record.Price
		}
		if temp.Low > record.Price || temp.Low == 0 {
			temp.Low = record.Price
		}
		result[record.StockCode] = temp
	}

	for _, indexMember := range indexMembers {
		temp := Summary{}
		if val, ok := result[indexMember.StockCode]; ok {
			temp = val
		}
		temp.IndexCode = append(temp.IndexCode, indexMember.IndexCode)
		temp.StockCode = indexMember.StockCode
		result[indexMember.StockCode] = temp
	}

	return result
}

var (
	recs = []ChangeRecord{
		{
			StockCode: "BBCA",
			Price:     8783,
			Quantity:  0,
		},
		{
			StockCode: "BBRI",
			Price:     3233,
			Quantity:  0,
		},
		{
			StockCode: "ASII",
			Price:     1223,
			Quantity:  0,
		},
		{
			StockCode: "GOTO",
			Price:     321,
			Quantity:  0,
		},

		{
			StockCode: "BBCA",
			Price:     8780,
			Quantity:  1,
		},
		{
			StockCode: "BBRI",
			Price:     3230,
			Quantity:  1,
		},
		{
			StockCode: "ASII",
			Price:     1220,
			Quantity:  1,
		},
		{
			StockCode: "GOTO",
			Price:     320,
			Quantity:  1,
		},

		{
			StockCode: "BBCA",
			Price:     8800,
			Quantity:  1,
		},
		{
			StockCode: "BBRI",
			Price:     3300,
			Quantity:  1,
		},
		{
			StockCode: "ASII",
			Price:     1300,
			Quantity:  1,
		},
		{
			StockCode: "GOTO",
			Price:     330,
			Quantity:  1,
		},

		{
			StockCode: "BBCA",
			Price:     8600,
			Quantity:  1,
		},
		{
			StockCode: "BBRI",
			Price:     3100,
			Quantity:  1,
		},
		{
			StockCode: "ASII",
			Price:     1100,
			Quantity:  1,
		},
		{
			StockCode: "GOTO",
			Price:     310,
			Quantity:  1,
		},

		{
			StockCode: "BBCA",
			Price:     8785,
			Quantity:  1,
		},
		{
			StockCode: "BBRI",
			Price:     3235,
			Quantity:  1,
		},
		{
			StockCode: "ASII",
			Price:     1225,
			Quantity:  1,
		},
		{
			StockCode: "GOTO",
			Price:     325,
			Quantity:  1,
		},
	}
	indexes = []IndexMember{
		{
			StockCode: "BBCA",
			IndexCode: "IHSG",
		},
		{
			StockCode: "BBRI",
			IndexCode: "IHSG",
		},
		{
			StockCode: "ASII",
			IndexCode: "IHSG",
		},
		{
			StockCode: "GOTO",
			IndexCode: "IHSG",
		},
		{
			StockCode: "BBCA",
			IndexCode: "LQ45",
		},
		{
			StockCode: "BBRI",
			IndexCode: "LQ45",
		},
		{
			StockCode: "ASII",
			IndexCode: "LQ45",
		},
		{
			StockCode: "GOTO",
			IndexCode: "LQ45",
		},
		{
			StockCode: "BBCA",
			IndexCode: "KOMPAS100",
		},
		{
			StockCode: "BBRI",
			IndexCode: "KOMPAS100",
		},
	}
)

func Main() {
	got := ohlc(recs, indexes)
	for _, v := range got {
		jsonValue, err := json.Marshal(v)
		if err != nil {
			fmt.Println("error: ", err.Error())
		}
		fmt.Println("summary: ", string(jsonValue))
	}
}
