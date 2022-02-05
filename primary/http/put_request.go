package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://127.0.0.1:8081/api/cart/selectAll"
	req, _ := http.NewRequest("PUT", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "483669576316362752")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
