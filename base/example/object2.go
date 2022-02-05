package main

import "fmt"

type User9 struct {
	name string
	age  int
}

//方法
func (user User9) getUser() {
	fmt.Printf("方法->用户姓名:%s 年龄:%d\n", user.name, user.age)
}

//函数
func getUser(user User9) {
	fmt.Printf("函数->用户姓名:%s 年龄:%d\n", user.name, user.age)
}

func main() {
	user := User9{name: "刘备", age: 60}
	user.getUser()
	getUser(user)
}
