package main

import (
	"fmt"
)

func main() {
	srcString := "A"
	destString := "a"

	if srcString == destString {
		fmt.Println("Equals")
	} else {
		fmt.Println("Not Equals")
	}
}
