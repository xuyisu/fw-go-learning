package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mutex := sync.Mutex{}
	for i := 0; i < 10; i++ {
		go func(i int) {
			mutex.Lock()
			defer mutex.Unlock()
			fmt.Println("i :=", i)
			time.Sleep(time.Second)
		}(i)
	}

	time.Sleep(15 * time.Second)
	fmt.Println("程序结束")
}
