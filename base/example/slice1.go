package main

import "fmt"

//切片

//切片是可索引的，并且可以由 len() 方法获取长度。
//切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

func main() {
	var numbers = make([]int, 3, 5)

	printSlice1(numbers)

	var numbers1 []int

	printSlice1(numbers1)
}

func printSlice1(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
