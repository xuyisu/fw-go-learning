package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	waitGroup := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)

		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println("current value: ", i)
			waitGroup.Done()
		}(i)
	}

	for {
		fmt.Println("start wait.")
		waitGroup.Wait()
		fmt.Println("wait finish.")
		break
	}

	time.Sleep(time.Second)
	fmt.Println("程序结束")
}
