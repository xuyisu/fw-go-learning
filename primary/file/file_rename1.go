package main

import (
	"fmt"
	"os"
)

func main() {
	//创建一个名为“fw_test.txt”的空文件
	file, err := os.Create("./fw_rename.txt") // 如果文件已存在，会将文件清空。
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
	err = os.Rename("./fw_rename.txt", "./fw_rename_new.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
}
