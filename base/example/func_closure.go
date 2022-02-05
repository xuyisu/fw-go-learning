package main

import "fmt"

//闭包  可以累加
func countByInt() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

//求和  非闭包  不可以累加
func countByInt2(i int) int {
	sum := 0
	sum += i
	return sum

}

func main() {
	count := countByInt()
	for i := 0; i <= 10; i++ {
		fmt.Print("i= ", i)
		fmt.Print(" sum1=", count(i))
		fmt.Println(" sum2=", countByInt2(i))
	}
}
