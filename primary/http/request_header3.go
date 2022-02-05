package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//设置Header信息即返回json 数据demo

type User1 struct {
	ID   string `json:"id"`
	NAME string `json:"name"`
	PWD  string `json:"pwd"`
}

func user1(w http.ResponseWriter, r *http.Request) {
	//声明
	var user1 User1
	user1.ID = "001"
	user1.NAME = "test"
	user1.PWD = "123456"
	message, err := json.Marshal(user1)
	if err != nil {
		fmt.Printf("json.Marshal err:%v", err)
		return
	}
	fmt.Println(user1)
	//设置为application/json
	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}

func main() {
	http.HandleFunc("/", user1)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
