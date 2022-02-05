package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	p := &sync.Pool{
		New: func() interface{} {
			return -1
		},
	}

	for i := 0; i < 5; i++ {
		go func(idx int) {
			p.Put(idx)
		}(i)
	}

	//取出来的对象是无序的
	for i := 0; i < 10; i++ {
		go func() {
			val := p.Get()
			fmt.Println("Get val :", val)
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("程序结束")
}
