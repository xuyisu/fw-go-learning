package main

import (
	"encoding/json"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger5 *zap.Logger

//请求体
type SysUserZap5 struct {
	Phone    string `json:"phone"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func init() {
	//获取编码器,NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder //指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	//文件writeSyncer
	fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./info.log", //日志文件存放目录
		MaxSize:    10,           //文件大小限制,单位MB
		MaxBackups: 50,           //最大保留日志文件数量
		MaxAge:     1,            //日志文件保留天数
		Compress:   false,        //是否压缩处理
	})
	fileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(fileWriteSyncer, zapcore.AddSync(os.Stdout)), zapcore.InfoLevel) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志,InfoLevel 记录出DEBUG 的级别

	logger5 = zap.New(fileCore, zap.AddCaller()) //AddCaller()为显示调用函数信息(包括文件名和行号)
	defer logger5.Sync()
}

func main() {
	userReq := SysUserZap5{
		Phone:    "15788888888",
		UserName: "关羽",
		Email:    "guanyu@go.com",
	}
	marshal, _ := json.Marshal(userReq)

	req := string(marshal)
	logger5.Debug(req)
	logger5.Info(req)
	logger5.Error(req)
	logger5.Warn(req)
}
