package main

import (
	"net/http"
	"strconv"
	"webapi1009/data"

	"github.com/gin-gonic/gin"
)

func handleGetAllEmployees(c *gin.Context) {
	emps := data.GetAllEmployees()
	c.IndentedJSON(http.StatusOK, emps)
}

func handleGetOneEmployee(c *gin.Context) {
	id := c.Param("id")
	numId, err := strconv.Atoi(id)
	if err == nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	employee := data.GetEmployeeById(numId)

	if employee == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
	} else {
		c.IndentedJSON(http.StatusOK, employee)
	}
}

func handleCreateEmployee(c *gin.Context) {
	var employee data.Employee
	if err := c.BindJSON(&employee); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	data.CreateNewEmployee(employee)
	c.JSON(http.StatusCreated, employee)
}

func main() {
	data.Init()

	r := gin.Default()

	r.GET("/api/employees", handleGetAllEmployees)
	r.GET("/api/employees/:id", handleGetOneEmployee)
	r.POST("/api/employees", handleCreateEmployee)
	/*r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  "http.StatusOK",
		})
	}) */
	r.Run(":8085")
}
