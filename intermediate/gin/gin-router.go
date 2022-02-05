package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// 创建v1组
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginGin)
	}
	// 创建v2组
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginGin2)
	}
	router.Run(":8080")

}

func loginGin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello LoginV1!",
	})
}

func loginGin2(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello LoginV2!",
	})
}
