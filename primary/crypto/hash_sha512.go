package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	cry := "hello go"
	hash := sha512.New()
	hash.Write([]byte(cry))
	sum := hash.Sum(nil)
	sha512 := fmt.Sprintf("%x", sum)
	fmt.Println(sha512)
}
