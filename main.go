package main

import (
	"bysykkel/api"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	client := api.NewAPIClient()
	//data, _ := api.GetStationData(client)
	data, err := api.GetStatusData(client)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Data: %v\n", data)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
