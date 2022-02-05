package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建基于cookie的存储引擎，myTestSessionId参数是用于加密的密钥
	var store = cookie.NewStore([]byte("myTestSessionId"))

	r := gin.Default()

	// 设置session中间件，参数GoSessionId，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("GoSessionId", store))

	r.GET("/login", func(c *gin.Context) {
		// 初始化session对象
		session := sessions.Default(c)
		id := "1234564564564564564"
		session.Set("authennticated", id)
		session.Save()
		c.JSON(200, gin.H{"message": "用户登录"})
	})
	r.GET("/info", func(c *gin.Context) {
		// 初始化session对象
		session := sessions.Default(c)
		auth := session.Get("authennticated")
		if auth == nil {
			c.JSON(500, gin.H{"message": "没有权限"})
		} else {
			c.JSON(200, gin.H{"message": "获取用户信息成功"})
		}
	})
	r.GET("/logout", func(c *gin.Context) {
		// 初始化session对象
		session := sessions.Default(c)
		session.Delete("authennticated")
		session.Save()
		c.JSON(200, gin.H{"message": "退出成功"})
	})
	r.Run(":8081")
}
