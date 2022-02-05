package main

import "fmt"

func main() {
	deferMethod()

}

func deferMethod() {
	defer func1()
	defer func2()
	defer func3()
}
func func1() {
	fmt.Println("a")
}
func func2() {
	fmt.Println("b")
}
func func3() {
	fmt.Println("c")
}
