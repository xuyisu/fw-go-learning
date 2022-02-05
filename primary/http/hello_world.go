package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello World")
}

func main() {
	server := &http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/", hello)
	server.ListenAndServe()

}
