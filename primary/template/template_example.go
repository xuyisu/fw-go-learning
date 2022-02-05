package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//如果取不到文件，配置一下 working directory
func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	// 1. 解析模板
	t, err := template.ParseFiles("template_example.tmpl", "template_ul.tmpl")
	if err != nil {
		fmt.Println("template parsefile failed, err:", err)
		return
	}

	// 2.渲染模板
	// 设置模板数据
	data := map[string]interface{}{
		"userName": "go 模板",
		"age":      19,
		"man":      true,
		"addr":     "上海",
		"Language": []string{"Go", "Python", "PHP", "JavaScript"},
	}
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", helloHandleFunc)
	http.ListenAndServe(":8081", nil)
}
