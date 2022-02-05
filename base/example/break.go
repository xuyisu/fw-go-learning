package main

import "fmt"

//break 语法
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
				break OuterLoop
			}
		}
	}
}
