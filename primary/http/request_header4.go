package main

import (
	"fmt"
	"net/http"
)

//设置Header401
func noAuth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
	fmt.Fprintln(w, "未授权，认证后才能访问该接口！")
}

func main() {
	http.HandleFunc("/noAuth", noAuth)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
