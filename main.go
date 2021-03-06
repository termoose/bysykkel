package main

import (
	"bysykkel/data"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		d, _ := data.GetFrontendData()

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"data": d,
		})
	})

	router.GET("/map/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		lat, long, err := data.GetLocation(id)

		if err != nil {
			c.Redirect(http.StatusMovedPermanently, "/")
		}

		c.HTML(http.StatusOK, "map.tmpl", gin.H{
			"lat": lat,
			"long": long,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(fmt.Sprintf(":%s", port))
}
