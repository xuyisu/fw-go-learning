package main

import (
	"encoding/json"
	"log"
)

//请求体
type SysUserReq struct {
	Phone    string `json:"phone"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func main() {

	userReq := SysUserReq{
		Phone:    "15788889999",
		UserName: "张飞",
		Email:    "zhangfei@go.com",
	}
	marshal, _ := json.Marshal(userReq)
	log.Print("Print user req=", string(marshal), "\n")
	log.Println("Println user req=", string(marshal))
	log.Printf("Printf user req=%v\n", string(marshal))
}
