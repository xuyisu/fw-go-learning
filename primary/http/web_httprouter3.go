package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type HostMap map[string]http.Handler

func (hs HostMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//根据域名获取对应的Handler路由，然后调用处理（分发机制）
	if handler := hs[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "Forbidden", 403)
	}
}

func main() {
	userRouter := httprouter.New()
	userRouter.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("sub1 test"))
	})

	dataRouter := httprouter.New()
	dataRouter.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("sub2 test"))
	})

	//分别用于处理不同的二级域名
	hs := make(HostMap)
	hs["sub1.localhost:8081"] = userRouter
	hs["sub2.localhost:8081"] = dataRouter

	log.Fatal(http.ListenAndServe(":8081", hs))
}
