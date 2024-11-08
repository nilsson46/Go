package main

import (
	"net/http"
	"strconv"
	"webapi1009/data"

	"github.com/gin-gonic/gin"
)

func handleStartPage(c *gin.Context) {
	c.String(http.StatusOK, "text/html; charset=utf-8", []byte("<h1>Start page</h1>"))
}

type PageView struct {
	Title  string
	Rubrik string
}

func handleAboutPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", &PageView{Title: "Title about html", Rubrik: "About page"})
}
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

func handleUpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	numId, _ := strconv.Atoi(id)
	employeeFromDB := data.GetEmployeeById(numId)
	if employeeFromDB == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	} else {
		var employeeJson data.Employee
		if err := c.BindJSON(&employeeJson); err != nil {
			return
		}

		employeeJson.Id = numId
		data.UpdateEmployee(employeeJson)
		c.IndentedJSON(http.StatusOK, employeeJson)
	}
}

func handleDeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	numId, _ := strconv.Atoi(id)
	employeeFromDB := data.GetEmployeeById(numId)
	if employeeFromDB == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	} else {
		data.DeleteEmployee(*employeeFromDB)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Employee deleted"})
	}
}

var config Config

func main() {
	readConfig(&config)
	data.Init(config.Database.File,
		config.Database.Server,
		config.Database.Database,
		config.Database.Username,
		config.Database.Password,
		config.Database.Port)

	r := gin.Default()

	r.LoadHTMLGlob("templates/**")
	r.GET("/", handleStartPage)

	r.GET("/about", handleAboutPage)

	r.GET("/api/employees", handleGetAllEmployees)
	r.GET("/api/employees/:id", handleGetOneEmployee)

	r.PUT("/api/employees/:id", handleUpdateEmployee)
	// Put = replace
	//Patch == update a few fields
	r.POST("/api/employees", handleCreateEmployee)
	/*r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  "http.StatusOK",
		})
	}) */

	r.DELETE("/api/employees/:id", handleDeleteEmployee)
	r.Run(":8085")
}
