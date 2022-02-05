package main

import "fmt"

func main() {
	// 创建一个3个元素缓冲大小的整型通道
	ch := make(chan int, 4)
	// 查看当前通道的大小
	fmt.Println(len(ch))
	// 发送3个整型元素到通道
	ch <- 11
	ch <- 12
	ch <- 13
	ch <- 14
	// 查看当前通道的大小
	fmt.Println(len(ch))
}
