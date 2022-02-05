package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	cry := "hello go"
	hash := sha1.New()
	hash.Write([]byte(cry))
	sum := hash.Sum(nil)
	sha1 := fmt.Sprintf("%x", sum)
	fmt.Println(sha1)
}
