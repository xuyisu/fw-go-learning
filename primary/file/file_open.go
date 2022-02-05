package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("fw.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)
	// 关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
}
