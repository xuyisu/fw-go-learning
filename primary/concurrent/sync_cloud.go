package main

import (
	"fmt"
	"sync"
	"time"
)

//控制读线程的阻塞
var done = false

func read(name string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		//在没写入的时候先阻塞
		c.Wait()
	}
	fmt.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	fmt.Println(name, "starts write")
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	c.L.Unlock()
	fmt.Println(name, "wakes all")
	// 唤醒所有等待的线程
	c.Broadcast()
}

func main() {

	// 使用互斥锁创建一个cond
	cond := sync.NewCond(&sync.Mutex{})

	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)

	time.Sleep(time.Second * 3)
	fmt.Println("程序结束")
}
