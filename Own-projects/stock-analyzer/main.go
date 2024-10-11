package stockanalyzer

import (
	"net/http"
	"stock/data"
	"strconv"

	"github.com/gin-gonic/gin"
)

func handleGetAllStocks(c *gin.Context) {
	stocks := data.GetAllStocks()
	c.IndentedJSON(http.StatusOK, stocks)
}

func handleGetStockById(c *gin.Context) {
	id := c.Param("id")
	numId, err := strconv.Atoi(id)
	if err == nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	stock := data.GetStockById(numId)

	if stock == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "stock not found"})
	} else {
		c.IndentedJSON(http.StatusOK, stock)
	}
}

func handleCreateStock(c *gin.Context) {
	var stock data.StockData
	if err := c.BindJSON(&stock); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	data.CreateNewStock(stock)
	c.JSON(http.StatusCreated, stock)
}

func main() {
	data.Init()

	r := gin.Default()

	r.GET("/api/stocks", handleGetAllStocks)
	r.GET("/api/stock/:id", handleGetStockById)
	r.POST("/api/stocks", handleCreateStock)
	/*r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  "http.StatusOK",
		})
	}) */
	r.Run(":8085")
}
