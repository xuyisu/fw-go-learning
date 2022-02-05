package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://127.0.0.1:8081/api/cart/1332985086563454976"
	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Authorization", "483669576316362752")
	req.Header.Add("Content-Type", "application/json")
	//如果此时不需要知道返回的错误值,如此则忽略了error变量
	res, _ := http.DefaultClient.Do(req)
	// defer 延迟执行语句常用于释放资源
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
