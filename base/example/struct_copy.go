package main

import "fmt"

//结构体

type Books3 struct {
	title   string
	author  string
	subject string
	book_id int
}

//浅拷贝与深拷贝
func main() {

	// 创建一个新的结构体
	book := Books3{"Go 语言", "www.go.com", "Go 语言教程", 101}

	fmt.Println("拷贝前：", book)
	book1 := book
	book1.title = "Go 语言深拷贝"
	fmt.Println("深拷贝后：", book1)
	fmt.Println("深拷贝后原始值：", book)
	book2 := &book
	book2.title = "Go 语言浅拷贝"
	fmt.Println("浅拷贝后：", *book2)
	fmt.Println("浅拷贝后原始值：", book)
}
