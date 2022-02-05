package main

import "fmt"

// 遍历切片中每个元素，通过给定的函数访问元素
func printValue(list []int, f func(int)) {
	for _, value := range list {
		f(value)
	}
}

func main() {
	sli := []int{1, 2, 3, 4, 5}
	// 使用匿名函数打印切片的内容
	printValue(sli, func(value int) {
		fmt.Println(value)
	})
}
