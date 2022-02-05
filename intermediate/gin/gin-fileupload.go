package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"path"
	"strconv"
	"time"
)

func main() {
	router := gin.Default()
	// 设置文件上传大小限制，默认是32m
	router.MaxMultipartMemory = 5 << 20 // 64 MiB

	router.POST("/upload", func(c *gin.Context) {
		// 获取上传文件，返回的是multipart.FileHeader对象，
		// 代表一个文件，里面包含了文件名之类的详细信息
		// file是表单字段名字
		file, _ := c.FormFile("file")
		// 打印上传的文件名
		log.Println(file.Filename)

		//时间戳纳秒实现的自定义文件名
		newFilename := strconv.Itoa(int(time.Now().UnixNano())) + path.Ext(file.Filename)

		c.SaveUploadedFile(file, "./uploads/"+newFilename)

		c.JSON(200, gin.H{
			"message": "upload success",
		})
	})
	router.Run(":8080")
}
