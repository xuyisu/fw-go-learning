package main

import (
	"fmt"
	"net/http"
)

//手动设置请求的Header 值
func Header(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	fmt.Fprintln(w, "set Content-Type ok")
}

func main() {
	http.HandleFunc("/header", Header)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
