package data

var stocks []StockData

func GetAllStocks() []StockData {
	return stocks
}

func CreateNewStock(stock StockData) {
	stocks = append(stocks, stock)
}

func GetStockById(id int) *StockData {
	for _, stock := range stocks {
		if stock.Id == id {
			return &stock
		}
	}
	return nil
}

func Init() {
	stocks = append(stocks, StockData{Id: 1, Stockname: "Google", LatestStockPrice: 1000, Date: "2020-10-10"})
	stocks = append(stocks, StockData{Id: 2, Stockname: "Apple", LatestStockPrice: 1500, Date: "2020-10-11"})
	stocks = append(stocks, StockData{Id: 3, Stockname: "Amazon", LatestStockPrice: 2000, Date: "2020-10-12"})
}
