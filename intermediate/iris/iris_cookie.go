package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

func main() {

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(logger.New())
	//设置cookie
	app.Get("/cookie", func(ctx iris.Context) {
		ctx.SetCookieKV("fw", "hello go")

		ctx.JSON(iris.Map{"message": "success"})
	})
	//获取cookie
	app.Get("/getCookie", func(ctx iris.Context) {

		cookie := ctx.GetCookie("fw")
		ctx.JSON(iris.Map{"message": cookie})
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
