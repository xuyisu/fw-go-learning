package main

import (
	"encoding/json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger4 *zap.Logger

//请求体
type SysUserZap4 struct {
	Phone    string `json:"phone"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func init() {
	writeSaucer, _ := os.Create("./info.log")                         //日志文件存放目录
	encoderConfig := zap.NewProductionEncoderConfig()                 //指定时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder             //指定时间编码器
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder           //在日志文件中使用大写字母记录日志级别
	encoder := zapcore.NewConsoleEncoder(encoderConfig)               //获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	core := zapcore.NewCore(encoder, writeSaucer, zapcore.DebugLevel) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	logger4 = zap.New(core, zap.AddCaller())                          //AddCaller()为显示调用函数信息(包括文件名和行号)
	defer logger4.Sync()
}

func main() {

	userReq := SysUserZap4{
		Phone:    "15788888888",
		UserName: "关羽",
		Email:    "guanyu@go.com",
	}
	marshal, _ := json.Marshal(userReq)

	req := string(marshal)
	logger4.Debug(req)
	logger4.Info(req)
	logger4.Error(req)
	logger4.Warn(req)

}
