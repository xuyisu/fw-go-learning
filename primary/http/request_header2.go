package main

import (
	"fmt"
	"net/http"
)

//设置请求重定向
func Redirect(w http.ResponseWriter, r *http.Request) {
	// 设置一个 301 重定向
	w.Header().Set("Location", "https://www.baidu.com")
	w.WriteHeader(301)
}

func main() {
	http.HandleFunc("/redirect", Redirect)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
