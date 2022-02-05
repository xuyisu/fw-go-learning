package main

import (
	"fmt"
	"os"
)

func main() {
	// 读写方式打开文件
	fp, err := os.OpenFile("fw.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fp)

	defer fp.Close() //关闭文件，释放资源。
}
