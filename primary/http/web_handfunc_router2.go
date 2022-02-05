package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

type welcomeHandler struct {
	Name string
}

func (h welcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi, %s", h.Name)
}

func main() {
	mux := http.NewServeMux()
	// 注册处理器函数
	mux.HandleFunc("/hello", hello2Handler)

	// 注册处理器
	mux.Handle("/hello/web", welcomeHandler{Name: "Hello Go Web"})

	server := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
