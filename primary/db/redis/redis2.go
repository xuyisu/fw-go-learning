package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func main() {
	//构建连接池
	pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   1024,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379", redis.DialDatabase(0), redis.DialPassword("123456"))
		},
	}

	c := pool.Get()
	defer c.Close()

	//设置数据
	//setValue2(c)
	//获取数据
	//getValue2(c)
	//批量设置数据
	//mSetValue2(c)
	//批量获取
	//mGetValue2(c)
	//hash 的设置
	//hSetValue2(c)
	//hash 的获取
	//hGetValue2(c)
	//设置过期时间
	//expire2(c)
	//队列
	//queueSet2(c)
	//队列取出
	//queueLPop2(c)
	//队列长度
	//queueLen2(c)
	//设置数据
	setNum2(c)
	//incr 数据
	incr2(c)
	//decr 数据
	decr2(c)
}

func queueLen2(c redis.Conn) {
	res, err := redis.Int(c.Do("llen", "fw:queue"))
	if err != nil {
		fmt.Println("llen error: ", err)
		return
	}
	fmt.Println(res)
}

func incr2(c redis.Conn) {
	res, err := redis.Int(c.Do("incr", "fw:yi:num"))
	if err != nil {
		fmt.Println("incr error: ", err)
		return
	}
	fmt.Println(res)
}

func decr2(c redis.Conn) {
	res, err := redis.Int(c.Do("decr", "fw:yi:num"))
	if err != nil {
		fmt.Println("llen error: ", err)
		return
	}
	fmt.Println(res)
}

func queueSet2(c redis.Conn) {
	_, err := c.Do("lpush", "fw:queue", "go", "java", 9)
	if err != nil {
		fmt.Println("lpush error: ", err)
		return
	}
}

func queueLPop2(c redis.Conn) {
	for {
		r, err := redis.String(c.Do("lpop", "fw:queue"))
		if err != nil {
			fmt.Println("lpop error: ", err)
			break
		}
		fmt.Println(r)
	}
}

func expire2(c redis.Conn) {
	_, err := c.Do("expire", "com:yi:go", 100000)
	if err != nil {
		fmt.Println("expire error: ", err)
		return
	}
}

func hGetValue2(c redis.Conn) {
	res, err := redis.String(c.Do("HGet", "fw:languages", "language"))
	if err != nil {
		fmt.Println("hget error: ", err)
		return
	}
	fmt.Println(res)
}

func hSetValue2(c redis.Conn) {
	_, err := c.Do("HSet", "fw:languages", "language", "go")
	if err != nil {
		fmt.Println("hset error: ", err)
		return
	}
}

func mGetValue2(c redis.Conn) {
	res, err := redis.Strings(c.Do("MGet", "fw:name", "fw:phone"))
	if err != nil {
		fmt.Println("MGet error: ", err)
		return
	}
	fmt.Println(res)
}

func mSetValue2(c redis.Conn) {
	_, err := c.Do("MSet", "fw:name", "fw", "fw:phone", "13666777888")
	if err != nil {
		fmt.Println("MSet error: ", err)
		return
	}
}

func getValue2(c redis.Conn) {
	res, err := redis.String(c.Do("Get", "com:yi:go"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func setValue2(c redis.Conn) {
	_, err := c.Do("Set", "com:yi:go", "刘备")
	if err != nil {
		fmt.Println(err)
		return
	}
}
func setNum2(c redis.Conn) {
	_, err := c.Do("Set", "fw:yi:num", 20)
	if err != nil {
		fmt.Println(err)
		return
	}
}
