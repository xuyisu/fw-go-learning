package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"io"
	"os"
	"path"
	"strconv"
	"time"
)

const maxSize = 5 << 20 // 5MB
func main() {

	app := iris.New()
	app.Use(logger.New())
	app.Post("/upload", iris.LimitRequestBodySize(maxSize+1<<20), func(ctx iris.Context) {
		// Get the file from the request.
		file, info, err := ctx.FormFile("file")
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}

		defer file.Close()

		//时间戳纳秒实现的自定义文件名
		newFilename := strconv.Itoa(int(time.Now().UnixNano())) + path.Ext(info.Filename)

		//必须先在当前目录建一个uploads 文件夹
		out, err := os.OpenFile("./uploads/"+newFilename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}
		defer out.Close()

		io.Copy(out, file)
	})
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
