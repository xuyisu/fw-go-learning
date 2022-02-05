package main

import "fmt"

//匿名结构体
func main() {

	// 创建一个新的结构体
	book := struct {
		title   string
		author  string
		subject string
		book_id int
	}{"Go 语言", "www.go.com", "Go 语言教程", 101}

	fmt.Println("result：", book)

}
