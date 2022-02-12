package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

func main() {

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(logger.New())
	//设置header
	app.Get("/header", func(ctx iris.Context) {
		ctx.Header("fw", "go header")

		ctx.JSON(iris.Map{"message": "success"})
	})
	//获取header
	app.Get("/getHeader", func(ctx iris.Context) {

		header := ctx.GetHeader("fw")
		ctx.JSON(iris.Map{"message": header})
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
