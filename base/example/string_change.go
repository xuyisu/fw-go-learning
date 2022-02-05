package main

import "fmt"

func main() {

	str := "hello world"
	change1(str)
	change2(str)
}

func change1(str string) {
	bytes := []byte(str)
	bytes[5] = ','
	fmt.Printf("byte str=%s\n", str)
	fmt.Printf("byte byte=%s\n", bytes)
}

func change2(str string) {
	runes := []rune(str)
	runes[5] = ','
	fmt.Println("rune str=", str)
	fmt.Println("rune runes=", string(runes))
}
