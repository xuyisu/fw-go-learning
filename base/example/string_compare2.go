package main

import (
	"fmt"
	"strings"
)

func main() {
	srcString := "A"
	destString := "a"

	if strings.Compare(strings.ToLower(srcString), strings.ToLower(destString)) == 0 {
		fmt.Println("Equals")
	} else {
		fmt.Println("Not Equals")
	}
}
