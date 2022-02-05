package main

import (
	"encoding/json"
	"go.uber.org/zap"
)

var sugaredLogger *zap.SugaredLogger

//请求体
type SysUserZap6 struct {
	Phone    string `json:"phone"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func init() {
	logger, _ := zap.NewDevelopment() //这里可以换成zap.NewExample()或zap.NewProduction()
	//调用Sugar()即可转sugaredLogger
	sugaredLogger = logger.Sugar()
	defer sugaredLogger.Sync()
}

func main() {
	userReq := SysUserZap6{
		Phone:    "15788888888",
		UserName: "关羽",
		Email:    "guanyu@go.com",
	}
	marshal, _ := json.Marshal(userReq)

	req := string(marshal)
	sugaredLogger.Debug(req)
	sugaredLogger.Info(req)
	sugaredLogger.Error(req)
	sugaredLogger.Warn(req)
}
