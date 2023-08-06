package main

import "github.com/gin-gonic/gin"

func IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.Run()
}
