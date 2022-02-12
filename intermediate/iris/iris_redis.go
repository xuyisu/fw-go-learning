package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
)

var (
	redisSessionId = "mySessionId"
	redisSession   = sessions.New(sessions.Config{Cookie: redisSessionId})
)

func infoRedis(ctx iris.Context) {
	var auth = redisSession.Start(ctx).GetString("authennticated")
	if auth == "" {
		ctx.JSON(iris.Map{"message": "没有权限"})
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	ctx.JSON(iris.Map{"message": "获取用户基本信息"})
}

func loginRedis(ctx iris.Context) {
	session := redisSession.Start(ctx)

	session.Set("authennticated", "211211651656156151")

	ctx.JSON(iris.Map{"message": "用户登录"})
}

func logoutRedis(ctx iris.Context) {
	session := redisSession.Start(ctx)
	// 撤销用户身份验证
	session.Destroy()
	ctx.JSON(iris.Map{"message": "退出成功"})
}

func main() {
	config := redis.Config{
		Network:   "tcp",
		Addr:      "127.0.0.1:6379",
		Clusters:  nil,
		Password:  "123456",
		Database:  "0",
		MaxActive: 10,
		Timeout:   redis.DefaultRedisTimeout,
		Prefix:    "",
		Delim:     redis.DefaultDelim,
		Driver:    redis.Redigo(),
	}
	database := redis.New(config)
	redisSession.UseDatabase(database)
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(logger.New())
	app.Get("/info", infoRedis)
	app.Get("/login", loginRedis)
	app.Get("/logout", logoutRedis)
	app.Run(iris.Addr(":8080"))

}
