package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello shoezoo gateway.")
	e := gin.Default()
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello shoezoo gateway from gin",
		})
	})
	e.Run()
}
