package main

import (
	"fmt"
	"strings"
)

func main() {
	srcString := "A"
	destString := "a"

	if strings.EqualFold(srcString, destString) == true {
		fmt.Println("Equals")
	} else {
		fmt.Println("Not Equals")
	}
}
