package main

import "fmt"

//defer 与return的执行顺序，defer  在return 之后调用的
var name string = "001"

func myFunc() string {
	defer func() {
		name = "002"
	}()

	fmt.Printf("myfunc 函数里的name：%s\n", name)
	return name
}

func main() {
	customName := myFunc()
	fmt.Printf("main 函数里的name: %s\n", name)
	fmt.Println("main 函数里的customName: ", customName)
}
