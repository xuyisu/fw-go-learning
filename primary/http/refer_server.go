package main

import "net/http"

type Refer struct {
	handler http.Handler
	refer   string
}

func (this *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == this.refer {
		this.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is handler"))
}

func hello2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
func main() {
	referer := &Refer{
		handler: http.HandlerFunc(myHandler),
		refer:   "localhost",
	}
	http.HandleFunc("/hello2", hello2)
	http.ListenAndServe(":8080", referer)
}
