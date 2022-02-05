package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
)

type User struct {
	Username, Password string
	RegTime            time.Time
}

func ShowTime(t time.Time, format string) string {
	return t.Format(format)
}

//如果取不到文件，配置一下 working directory
func showFunc(w http.ResponseWriter, r *http.Request) {

	htmlByte, err := ioutil.ReadFile("./template_msg.tmpl")
	if err != nil {
		fmt.Println("read html failed, err:", err)
		return
	}

	u := User{"admin", "123456", time.Now()}
	t, err := template.New("text").Funcs(template.FuncMap{"showtime": ShowTime}).
		Parse(string(htmlByte))
	if err != nil {
		panic(err)
	}
	t.Execute(w, u)
}

func main() {
	http.HandleFunc("/", showFunc)
	http.ListenAndServe(":8081", nil)
}
