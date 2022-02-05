package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	rwMutex := sync.RWMutex{}
	for i := 0; i < 10; i++ {
		go func(idx int) {
			rwMutex.RLock()
			defer rwMutex.RUnlock()
			fmt.Println("Read Mutex :", idx)
		}(i)

		go func(idx int) {
			rwMutex.Lock()
			defer rwMutex.Unlock()
			fmt.Println("Write Mutex :", idx)
			time.Sleep(time.Second)
		}(i)
	}

	time.Sleep(20 * time.Second)
	fmt.Println("程序结束")
}
