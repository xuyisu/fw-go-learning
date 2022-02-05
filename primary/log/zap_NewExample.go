package main

import (
	"encoding/json"
	"go.uber.org/zap"
)

var logger2 *zap.Logger

//请求体
type SysUserZap2 struct {
	Phone    string `json:"phone"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func init() {
	logger2 = zap.NewExample()
	defer logger2.Sync()
}

func main() {
	userReq := SysUserZap2{
		Phone:    "15788888888",
		UserName: "关羽",
		Email:    "guanyu@go.com",
	}
	marshal, _ := json.Marshal(userReq)

	req := string(marshal)
	logger2.Debug(req)
	logger2.Info(req)
	logger2.Error(req)
	logger2.Warn(req)

}
