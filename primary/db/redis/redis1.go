package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//redis.Dial 案例
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
	//设置数据
	//setValue(err, c)
	//获取数据
	//getValue(err, c)
	//批量设置数据
	//mSetValue(err, c)
	//批量获取
	//mGetValue(err, c)
	//hash 的设置
	//hSetValue(err, c)
	//hash 的获取
	//hGetValue(err, c)
	//设置过期时间
	//expire(err, c)
	//队列
	//queueSet(err, c)
	//队列取出
	//queueLPop(err, c)
	//队列长度
	//queueLen(err, c)
	//设置数据
	setNum(err, c)
	//incr 数据
	incr(err, c)
	//decr 数据
	decr(err, c)

}

func queueLen(err error, c redis.Conn) {
	res, err := redis.Int(c.Do("llen", "fw:queue"))
	if err != nil {
		fmt.Println("llen error: ", err)
		return
	}
	fmt.Println(res)
}

func incr(err error, c redis.Conn) {
	res, err := redis.Int(c.Do("incr", "fw:yi:num"))
	if err != nil {
		fmt.Println("incr error: ", err)
		return
	}
	fmt.Println(res)
}

func decr(err error, c redis.Conn) {
	res, err := redis.Int(c.Do("decr", "fw:yi:num"))
	if err != nil {
		fmt.Println("llen error: ", err)
		return
	}
	fmt.Println(res)
}

func queueSet(err error, c redis.Conn) {
	_, err = c.Do("lpush", "fw:queue", "go", "java", 9)
	if err != nil {
		fmt.Println("lpush error: ", err)
		return
	}
}

func queueLPop(err error, c redis.Conn) {
	for {
		r, err := redis.String(c.Do("lpop", "fw:queue"))
		if err != nil {
			fmt.Println("lpop error: ", err)
			break
		}
		fmt.Println(r)
	}
}

func expire(err error, c redis.Conn) {
	_, err = c.Do("expire", "com:yi:go", 100000)
	if err != nil {
		fmt.Println("expire error: ", err)
		return
	}
}

func hGetValue(err error, c redis.Conn) {
	res, err := redis.String(c.Do("HGet", "fw:languages", "language"))
	if err != nil {
		fmt.Println("hget error: ", err)
		return
	}
	fmt.Println(res)
}

func hSetValue(err error, c redis.Conn) {
	_, err = c.Do("HSet", "fw:languages", "language", "go")
	if err != nil {
		fmt.Println("hset error: ", err)
		return
	}
}

func mGetValue(err error, c redis.Conn) {
	res, err := redis.Strings(c.Do("MGet", "fw:name", "fw:phone"))
	if err != nil {
		fmt.Println("MGet error: ", err)
		return
	}
	fmt.Println(res)
}

func mSetValue(err error, c redis.Conn) {
	_, err = c.Do("MSet", "fw:name", "fw", "fw:phone", "13666777888")
	if err != nil {
		fmt.Println("MSet error: ", err)
		return
	}
}

func getValue(err error, c redis.Conn) {
	res, err := redis.String(c.Do("Get", "com:yi:go"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func setValue(err error, c redis.Conn) {
	_, err = c.Do("Set", "com:yi:go", "刘备")
	if err != nil {
		fmt.Println(err)
		return
	}
}
func setNum(err error, c redis.Conn) {
	_, err = c.Do("Set", "fw:yi:num", 20)
	if err != nil {
		fmt.Println(err)
		return
	}
}
