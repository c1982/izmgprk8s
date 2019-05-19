package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	host, _ := os.Hostname()

	g := gin.New()
	g.GET("/", func(c *gin.Context) {
		c.String(200, "hello gophers! Host: "+host)
	})

	g.Run(":8080")
}
