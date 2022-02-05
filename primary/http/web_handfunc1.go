package main

import (
	"fmt"
	"log"
	"net/http"
)

//http.HandleFunc  方式启动server 端
func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", hi)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
