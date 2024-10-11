package data

type StockData struct {
	Id               int
	Stockname        string
	LatestStockPrice int
	Date             string
}

type AnalysisResult struct {
	Stockname string
	Average   float64
	Max       int
	Min       int
	Trend     string
}

func NewStockData(id int, stockname string, latestStockPrice int, date string) *StockData {
	return &StockData{
		Id:               id,
		Stockname:        stockname,
		LatestStockPrice: latestStockPrice,
		Date:             date,
	}
}
