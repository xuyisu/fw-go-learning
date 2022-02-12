package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	//实例化echo对象。
	e := echo.New()

	//设置session数据保存目录
	sessionPath := "./session_data"

	//设置cookie加密秘钥, 可以随意设置
	var sessionKey = "myTestSessionId"

	//输出日志
	e.Use(middleware.Logger())
	//设置session中间件
	//这里使用的session中间件，session数据保存在指定的目录
	e.Use(session.Middleware(sessions.NewFilesystemStore(sessionPath, []byte(sessionKey))))

	e.GET("/login", func(c echo.Context) error {

		sess, _ := session.Get("GoSessionId", c)
		sess.Values["authennticated"] = "1234564564564564564"
		sess.Save(c.Request(), c.Response())
		var res = make(map[string]string)
		res["message"] = "登录成功"
		return c.JSON(http.StatusOK, res)
	})
	e.GET("/info", func(c echo.Context) error {
		sess, _ := session.Get("GoSessionId", c)
		id := sess.Values["authennticated"]
		var res = make(map[string]string)
		if id != nil {
			res["message"] = "成功查询用户信息"
		} else {
			res["message"] = "没有权限"
		}
		return c.JSON(http.StatusOK, res)
	})
	e.GET("/logout", func(c echo.Context) error {

		sess, _ := session.Get("GoSessionId", c)
		sess.Values["authennticated"] = nil
		sess.Save(c.Request(), c.Response())
		var res = make(map[string]string)
		res["message"] = "登出成功"
		return c.JSON(http.StatusOK, res)
	})

	e.Start(":8080")
}
