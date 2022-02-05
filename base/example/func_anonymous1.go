package main

import "fmt"

//匿名函数
func main() {
	// 定义匿名函数并赋值给f变量
	f := func(data int) {
		fmt.Println("closure", data)
	}
	// 此时f变量的类型是func(), 可以直接调用
	f(6)

	//直接声明并调用
	func(data int) {
		fmt.Println("closure-directly", data)
	}(8)
}
