package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 使用 io/ioutil.ReadFile 方法一次性将文件读取到内存中
	filePath := "fw.txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", string(content))
}
