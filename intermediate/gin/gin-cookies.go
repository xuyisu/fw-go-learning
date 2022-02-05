package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/cookie", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("fw", "hello go", 3600, "/", "localhost", false, true)
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.GET("/getCookie", func(c *gin.Context) {
		// 根据cookie名字读取cookie值
		data, err := c.Cookie("fw")
		if err != nil {
			fmt.Println("异常=", err)
		}
		// 直接返回cookie值
		c.JSON(200, gin.H{
			"message": data,
		})
	})

	r.Run()
}
