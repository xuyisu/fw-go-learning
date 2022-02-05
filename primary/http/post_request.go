package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://127.0.0.1:8081/api/user/login"
	body := "{\"userName\":\"admin\",\"password\":\"123456\"}"
	response, err := http.Post(url, "application/json",
		bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println("err", err)
	}
	b, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(b))
}
