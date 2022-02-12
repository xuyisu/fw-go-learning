package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	//实例化echo对象。
	e := echo.New()
	//输出日志
	e.Use(middleware.Logger())
	//注册一个Get请求, 路由地址为: /cookie  并且绑定一个控制器函数, 这里使用的是闭包函数。
	e.GET("/cookie", func(c echo.Context) error {

		var res = make(map[string]string)
		res["message"] = "success"

		//初始化cookie对象
		cookie := new(http.Cookie)
		cookie.Name = "fw"
		cookie.Value = "hello go"
		cookie.Path = "/"
		//cookie有效期为3600秒
		cookie.MaxAge = 3600
		cookie.HttpOnly = true
		cookie.Secure = false
		// 设置cookie
		c.SetCookie(cookie)
		return c.JSON(http.StatusOK, res)
	})

	//注册一个Get请求, 路由地址为: /hello  并且绑定一个控制器函数, 这里使用的是闭包函数。
	e.GET("/getCookie", func(c echo.Context) error {

		var res = make(map[string]string)
		//根据cookie名，获取cookie, cookie存在则返回http.Cookie结构体
		cookie, err := c.Cookie("fw")
		if err != nil {
			return err
		}
		res["message"] = cookie.Value
		return c.JSON(http.StatusOK, res)
	})

	//启动http server, 并监听8080端口，冒号（:）前面为空的意思就是绑定网卡所有Ip地址，本机支持的所有ip地址都可以访问。
	e.Start(":8080")
}
