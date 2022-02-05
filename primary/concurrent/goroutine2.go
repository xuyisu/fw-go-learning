package main

import (
	"fmt"
	"time"
)

func main() {
	go msg1("hello go")
	go msg2("hi go")
	time.Sleep(3 * time.Second)
}

func msg1(msg string) {
	for i := 0; i < 5; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("msg1:", msg)
	}
}

func msg2(msg string) {
	for i := 0; i < 5; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("msg2:", msg)
	}
}
