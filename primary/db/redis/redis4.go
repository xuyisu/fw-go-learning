package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	//设置db
	database := redis.DialDatabase(0)
	//设置密码
	password := redis.DialPassword("123456")
	c, err := redis.Dial("tcp", "localhost:6379", database, password)
	if err != nil {
		fmt.Println("conn redis failed, err:", err)
		return
	}
	defer c.Close()
	//开启事务
	c.Send("MULTI")
	c.Send("SET", "fw:go", "go")
	c.Send("SET", "fw:java", "java")
	//执行事务
	r, err := c.Do("EXEC")
	if err != nil {
		//取消事务
		c.Do("DISCARD")
	}
	fmt.Println(r)
}
