package main

import (
	"net/http"
)

//设置cookie
func cookieHandle(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:   "token",
		Value:  "AAC8A721-FF21-4B04-A350-AE1A243A92A0",
		MaxAge: 3600,
		Domain: "localhost",
		Path:   "/",
	}

	http.SetCookie(w, cookie)

	//在具体数据返回之前设置cookie，否则cookie种不上
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/", cookieHandle)
	http.ListenAndServe(":8085", nil)
}
