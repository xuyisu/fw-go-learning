package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//不设置随机种子，每次运行结果都一样
	rand.Seed(time.Now().UnixNano())
	//如果n<=0会panic  取值范围[0,n)
	fmt.Println("指定范围的伪随机int值:", rand.Intn(10))
	fmt.Println("指定范围的伪随机int32值:", rand.Int31n(100))
	fmt.Println("指定范围的伪随机int64值:", rand.Int63n(10000000))
	fmt.Println("伪随机int64数", rand.Int63())
	fmt.Println("伪随机int32数", rand.Int31())
	fmt.Println("伪随机int值", rand.Int())
	fmt.Println("伪随机uint32值", rand.Uint32())
	fmt.Println("伪随机uint64值", rand.Uint64())
	fmt.Println("伪随机float32值", rand.Float32())
	fmt.Println("伪随机float64值", rand.Float64())
	fmt.Println("随机数切片:", rand.Perm(5))
}
