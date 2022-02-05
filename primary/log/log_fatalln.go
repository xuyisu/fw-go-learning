package main

import (
	"encoding/json"
	"log"
)

//请求体
type SysUserReq2 struct {
	Phone    string `json:"phone"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func main() {
	userReq := SysUserReq2{
		Phone:    "15788888888",
		UserName: "关羽",
		Email:    "guanyu@go.com",
	}
	marshal, _ := json.Marshal(userReq)

	log.Fatal("Fatal user req=", string(marshal), "\n")
	log.Fatalln("Fatalln user req=", string(marshal))
	log.Fatalf("Fatalf user req=%v\n", string(marshal))
}
