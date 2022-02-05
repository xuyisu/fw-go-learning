package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	cry := "hello go"
	hash := sha256.New()
	hash.Write([]byte(cry))
	sum := hash.Sum(nil)
	sha256 := fmt.Sprintf("%x", sum)
	fmt.Println(sha256)
}
