package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type (
	OrderDetail2 struct {
		OrderId       string `json:"orderId"`
		ReceiveName   string `json:"receiveName"`
		ReceiveMobile string `json:"receiveMobile"`
		ReceiveAddr   string `json:"receiveAddr"`
	}

	Order2 struct {
		OrderNo     string       `json:"orderNo"`
		OrderDetail OrderDetail2 `json:"orderDetail"`
	}
)

func main() {
	orderDetail := OrderDetail2{
		OrderId:       "1010203035363231",
		ReceiveName:   "fw",
		ReceiveMobile: "15899998888",
		ReceiveAddr:   "安徽省XXX市",
	}

	order := Order2{
		OrderNo:     "210987",
		OrderDetail: orderDetail,
	}

	data, err := json.Marshal(&order)
	if err != nil {
		fmt.Printf("json.Marshal failed,err:%v", err)
		return
	}

	fmt.Printf("%s\n", string(data))

	file, _ := os.Create("json_write2.json")
	defer file.Close()
	file.Write(data)
}
