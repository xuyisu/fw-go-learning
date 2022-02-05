package main

import "fmt"

func main() {
	c := make(chan string)
	//开启一个匿名函数发送hello
	go func() { c <- "hello" }()
	//接收发送的数据
	msg := <-c
	fmt.Println(msg)
}
