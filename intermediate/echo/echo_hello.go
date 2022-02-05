package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	//实例化echo对象。
	e := echo.New()
	//注册一个Get请求, 路由地址为: /hello  并且绑定一个控制器函数, 这里使用的是闭包函数。
	e.GET("/hello", func(c echo.Context) error {

		var res = make(map[string]string)
		res["message"] = "hello echo"
		return c.JSON(http.StatusOK, res)
	})

	//启动http server, 并监听8080端口，冒号（:）前面为空的意思就是绑定网卡所有Ip地址，本机支持的所有ip地址都可以访问。
	e.Start(":8080")
}
