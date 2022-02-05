package main

import "fmt"

func main() {

	condition()

}

func one() {
	var res = "a"
	switch res {
	case "a":
		fmt.Println("结果为a")
	case "b":
		fmt.Println("结果为b")
	case "c":
		fmt.Println("结果为c")
	default:
		fmt.Println("结果为 default")
	}
}

func many() {
	var res = "a"
	switch res {
	case "a", "b":
		fmt.Println("结果为a或b")
	case "c":
		fmt.Println("结果为c")
	case "d":
		fmt.Println("结果为d")
	default:
		fmt.Println("结果为 default")
	}
}

func condition() {
	var res = "a"
	switch {
	case res == "a" || res == "b":
		fmt.Println("结果为a或b")
	default:
		fmt.Println("结果为 default")
	}
}
