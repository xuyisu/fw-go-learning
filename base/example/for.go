package main

import "fmt"

//循环
func main() {

	var i, j int
	for i = 2; i < 10; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d 是素数\n", i)
		}
	}

}
