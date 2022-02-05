package main

import "fmt"

//选取最大值
func main() {
	var a int = 100
	var b int = 200
	var ret int
	ret = max(a, b)
	fmt.Printf("最大值是 : %d\n", ret)
}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
