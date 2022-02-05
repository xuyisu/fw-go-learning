package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	cry := "hello go"
	hash := md5.New()
	hash.Write([]byte(cry))
	sum := hash.Sum(nil)
	md5 := fmt.Sprintf("%x", sum)
	fmt.Println(md5)
}
