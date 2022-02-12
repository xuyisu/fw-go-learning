package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

func main() {
	//实例化echo对象。
	e := echo.New()
	//输出日志
	e.Use(middleware.Logger())
	//注册一个Get请求, 路由地址为: /hello  并且绑定一个控制器函数, 这里使用的是闭包函数。
	e.POST("/upload", func(c echo.Context) error {

		var res = make(map[string]string)
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		//打开用户上传的文件
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		//时间戳纳秒实现的自定义文件名
		newFilename := strconv.Itoa(int(time.Now().UnixNano())) + path.Ext(file.Filename)

		//必须先在当前目录建一个uploads 文件夹
		out, err := os.OpenFile("./uploads/"+newFilename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		defer out.Close()

		io.Copy(out, src)
		res["message"] = "upload success"
		return c.JSON(http.StatusOK, res)
	})

	//启动http server, 并监听8080端口，冒号（:）前面为空的意思就是绑定网卡所有Ip地址，本机支持的所有ip地址都可以访问。
	e.Start(":8080")
}
