package main

import (
	"encoding/json"
	"log"
)

//请求体
type SysUserReq3 struct {
	Phone    string `json:"phone"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func main() {
	userReq := SysUserReq3{
		Phone:    "15788888888",
		UserName: "刘备",
		Email:    "liubei@go.com",
	}
	marshal, _ := json.Marshal(userReq)

	log.Panic("Panic user req=", string(marshal), "\n")
	log.Panicln("Panicln user req=", string(marshal))
	log.Panicf("Panicf user req=%v\n", string(marshal))
}
