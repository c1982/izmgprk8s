package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	g := gin.New()
	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Mother Gophers!")
	})

	g.Run(":8080")
}
