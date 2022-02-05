package main

import (
	"encoding/json"
	"go.uber.org/zap"
)

var logger3 *zap.Logger

//请求体
type SysUserZap3 struct {
	Phone    string `json:"phone"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func init() {
	logger3, _ = zap.NewDevelopment()
	defer logger3.Sync()
}

func main() {
	userReq := SysUserZap3{
		Phone:    "15788888888",
		UserName: "关羽",
		Email:    "guanyu@go.com",
	}
	marshal, _ := json.Marshal(userReq)

	req := string(marshal)
	logger3.Debug(req)
	logger3.Info(req)
	logger3.Error(req)
	logger3.Warn(req)

}
