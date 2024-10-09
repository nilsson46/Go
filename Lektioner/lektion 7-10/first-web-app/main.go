package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type PageView struct {
	Title       string
	CurrentUser string
	PageTitle   string
	Text        string
	Header      string
}

// handler
func start(c *gin.Context) {
	log.Println("Handling / request")
	c.HTML(200, "home", &PageView{Title: "Home", CurrentUser: "Guest", PageTitle: "Welcome", Text: "Welcome to the home page!", Header: "Welcome to My Website"})
}

func about(c *gin.Context) {
	log.Println("Handling /about request")
	c.String(200, "About")
}

func main() {
	log.Println("Starting application")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", start)
	router.GET("/about", about)
	err := router.Run(":8084")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
