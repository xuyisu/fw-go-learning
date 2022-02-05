package main

import "fmt"

//引用传递
func changeP(x, y *int) {
	fmt.Println("开始变更")
	tmp := *x
	*x = *y
	*y = tmp
}

func main() {
	x := 10
	y := 20
	fmt.Println("变更前x:", x)
	fmt.Println("变更前y:", y)
	changeP(&x, &y)
	fmt.Println("变更后x:", x)
	fmt.Println("变更后y:", y)

}
