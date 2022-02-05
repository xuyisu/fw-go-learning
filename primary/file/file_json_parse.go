package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//定义配置文件解析后的结构
type (
	OrderDetail struct {
		OrderId       string `json:"orderId"`
		ReceiveName   string `json:"receiveName"`
		ReceiveMobile string `json:"receiveMobile"`
		ReceiveAddr   string `json:"receiveAddr"`
	}

	Order struct {
		OrderNo     string      `json:"orderNo"`
		OrderDetail OrderDetail `json:"orderDetail"`
	}
)

func main() {
	v := Order{}
	data, err := ioutil.ReadFile("./json_parse.json")
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err)
	}
	json, _ := json.Marshal(v)
	fmt.Println("JSON data is :", string(json))
	fmt.Println("OrderNo is :", v.OrderNo)
	fmt.Println("OrderId is :", v.OrderDetail.OrderId)
	fmt.Println("ReceiveName is :", v.OrderDetail.ReceiveName)
	fmt.Println("ReceiveAddr is :", v.OrderDetail.ReceiveAddr)
	fmt.Println("ReceiveMobile is :", v.OrderDetail.ReceiveMobile)
}
