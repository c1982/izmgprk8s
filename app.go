package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String("hello gophers!")
	})

	r.Run(":8080")
}
