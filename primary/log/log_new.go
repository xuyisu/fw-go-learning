package main

import (
	"encoding/json"
	"log"
	"os"
)

//请求体
type SysUserReq4 struct {
	Phone    string `json:"phone"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

//更改日志级别
func main() {

	userReq := SysUserReq4{
		Phone:    "15788888888",
		UserName: "关羽",
		Email:    "guanyu@go.com",
	}
	marshal, _ := json.Marshal(userReq)

	fileName := "New.log"
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error")
	}
	//输出到文件
	//debugLog := log.New(logFile, "[Info]", log.Llongfile)
	//输出到控制台
	debugLog := log.New(os.Stdout, "[Info]", log.Llongfile)
	debugLog.Println(string(marshal))
}
