package main

import "fmt"

//函数返回多个值
func compute(a, b int) (int, int) {
	//return a,b  //返回10 20
	return b, a //返回 20 10
}

func main() {
	a, b := compute(10, 20)
	fmt.Println(a, b)
}
