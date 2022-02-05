package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/header", func(c *gin.Context) {
		c.Header("fw", "go header")
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.GET("/getHeader", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": c.GetHeader("fw"),
		})
	})

	r.Run()
}
