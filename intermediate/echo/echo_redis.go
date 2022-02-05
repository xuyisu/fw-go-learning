package main

import (
	"github.com/gin-contrib/sessions/redis"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	//实例化echo对象。
	e := echo.New()

	// 创建基于cookie的存储引擎，mySessionId参数是用于加密的密钥
	var storeRedis, _ = redis.NewStore(10, "tcp", "localhost:6379", "123456", []byte("passord"))

	//保存到Redis
	e.Use(session.Middleware(storeRedis))

	e.GET("/login", func(c echo.Context) error {

		sess, _ := session.Get("user_session", c)
		sess.Values["fw"] = "156156451148498498494"
		sess.Save(c.Request(), c.Response())
		var res = make(map[string]string)
		res["message"] = "登录成功"
		return c.JSON(http.StatusOK, res)
	})
	e.GET("/info", func(c echo.Context) error {
		sess, _ := session.Get("user_session", c)
		id := sess.Values["fw"]
		var res = make(map[string]string)
		if id != nil {
			res["message"] = "成功查询用户信息"
		} else {
			res["message"] = "没有权限"
		}
		return c.JSON(http.StatusOK, res)
	})
	e.GET("/logout", func(c echo.Context) error {

		sess, _ := session.Get("user_session", c)
		sess.Values["fw"] = nil
		sess.Save(c.Request(), c.Response())
		var res = make(map[string]string)
		res["message"] = "登出成功"
		return c.JSON(http.StatusOK, res)
	})

	e.Start(":8080")
}
