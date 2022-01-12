package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"message": "hello",
		})
	})
	r.Run()
}