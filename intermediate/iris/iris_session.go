package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var (
	testSessionId = "mySessionId"
	sess          = sessions.New(sessions.Config{Cookie: testSessionId})
)

func info(ctx iris.Context) {
	var auth = sess.Start(ctx).GetString("authennticated")
	if auth == "" {
		ctx.JSON(iris.Map{"message": "没有权限"})
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	ctx.JSON(iris.Map{"message": "获取用户基本信息"})
}

func login(ctx iris.Context) {
	session := sess.Start(ctx)

	session.Set("authennticated", "211211651656156151")

	ctx.JSON(iris.Map{"message": "用户登录"})
}

func logout(ctx iris.Context) {
	session := sess.Start(ctx)
	// 撤销用户身份验证
	session.Destroy()
	ctx.JSON(iris.Map{"message": "退出成功"})
}

func main() {
	app := iris.New()

	app.Get("/info", info)
	app.Get("/login", login)
	app.Get("/logout", logout)
	app.Run(iris.Addr(":8080"))

}
