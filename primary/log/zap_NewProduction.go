package main

import (
	"encoding/json"
	"go.uber.org/zap"
)

var logger *zap.Logger

//请求体
type SysUserZap struct {
	Phone    string `json:"phone"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func init() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
}

func main() {
	userReq := SysUserZap{
		Phone:    "15788888888",
		UserName: "关羽",
		Email:    "guanyu@go.com",
	}
	marshal, _ := json.Marshal(userReq)
	//debug  默认不输出
	req := string(marshal)
	logger.Debug(req)
	logger.Info(req)
	logger.Error(req)
	logger.Warn(req)

}
