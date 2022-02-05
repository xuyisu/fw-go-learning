package main

import "fmt"

//条件操作
func main() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)

	if a == b {
		fmt.Printf("第四行 - a 等于 b\n")
	} else {
		fmt.Printf("第五行 - a 不等于 b\n")
	}
}
