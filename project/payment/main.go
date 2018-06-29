package main

import (
	"github.com/gin-gonic/gin"
	"sillyhat/project/payment/readconfig"
)

func main() {
	readconfig.ReadConfigUtils()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":18002")
}