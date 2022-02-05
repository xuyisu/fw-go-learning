package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			c <- i
		}
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("程序结束")
}
