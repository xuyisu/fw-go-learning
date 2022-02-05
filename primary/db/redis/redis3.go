package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379", redis.DialDatabase(0), redis.DialPassword("123456"))
	if err != nil {
		fmt.Println("conn redis failed, err:", err)
		return
	}
	defer c.Close()
	//单个set
	setValue3(c)
	//hash set
	hsetValue3(c)
}

func hsetValue3(c redis.Conn) {
	c.Send("HSET", "languages", "language", "go", "desc", "hello go")
	c.Flush()
	receive, err := c.Receive()
	//receive 返回影响的field 数量
	fmt.Printf("receive:%v,err=%v\n", receive, err)
}

func setValue3(c redis.Conn) {
	c.Send("SET", "fw:yi:go", "刘备")
	c.Send("SET", "fw:yi:java", "张飞")
	c.Send("SET", "fw:yi:nodejs", "赵云")
	c.Flush()
	receive, err := c.Receive()
	fmt.Printf("receive:%v,err=%v\n", receive, err)
	receive1, err1 := c.Receive()
	fmt.Printf("receive:%v,err=%v\n", receive1, err1)
	receive2, err2 := c.Receive()
	fmt.Printf("receive:%v,err=%v\n", receive2, err2)
}
