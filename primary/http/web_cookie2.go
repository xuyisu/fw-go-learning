package main

import (
	"fmt"
	"net/http"
)

//设置cookie
func cookieGetHandle(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("fw")
	fmt.Printf("cookie:%#v, err:%v\n", c, err)
	//在具体数据返回之前设置cookie，否则cookie种不上
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/", cookieGetHandle)
	http.ListenAndServe(":8085", nil)
}
