package main

import "fmt"

//continue  语法
func main() {
	//外层循环的标签
OuterLoop:
	//双层循环
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			//使用 switch 进行数值分支判断
			switch i {
			case 1:
				//换行输出
				fmt.Println(i, j)
				//退出 OuterLoop 对应的循环之外
				continue OuterLoop
			case 2:
				fmt.Println(i, j)
				continue OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}

		}
	}
}
