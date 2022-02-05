package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/getUser", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("get 获取用户"))
	})
	router.POST("/user", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("post 用户信息"))
	})
	//精确匹配
	router.GET("/user/:name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("用户 name:" + p.ByName("name")))
	})
	//匹配所有
	router.GET("/info/*name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("接口info name:" + p.ByName("name")))
	})
	http.ListenAndServe(":8081", router)
}
