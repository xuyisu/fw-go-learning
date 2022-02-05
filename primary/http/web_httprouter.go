package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

//第三方组件设置路由
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Index"))
}

func Hello3(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Hello World"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello", Hello3)
	log.Fatal(http.ListenAndServe(":8081", router))
}
