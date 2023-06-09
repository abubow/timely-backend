package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
func main() {
	fmt.Println("Hello World!")
	r := gin.Default()
	r.GET("/hello", hello)
	r.Run()
}
