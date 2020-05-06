package main

import (
	"bysykkel/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	client := api.NewAPIClient()
	data, err := api.GetStationData(client)
	//data, err := api.GetStatusData(client)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Data: %v\n", data)

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "hello lol",
			"stations": data.Data.Stations,
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}
