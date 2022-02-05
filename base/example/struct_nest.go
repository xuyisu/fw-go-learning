package main

import "fmt"

type Books5 struct {
	title   string
	author  string
	subject string
	book_id int
}

type Company struct {
	name string
	book Books5
}

//结构体嵌套
func main() {

	book := Books5{"Go 语言", "www.go.com", "Go 语言教程", 101}

	company := Company{}
	company.name = "刘备公司"
	company.book = book
	fmt.Println("company：", company)

}
