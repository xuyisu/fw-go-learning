package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建文件
	//文件的创建，Create会根据传入的文件名创建文件，默认权限是0666//-rw-r--r--
	fp, err := os.Create("./fw.txt") // 如果文件已存在，会将文件清空。
	// defer延迟调用
	defer fp.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileInfo, err := os.Stat("./fw.txt")
	fileMode := fileInfo.Mode()
	fmt.Println(fileMode)
	os.Chmod("./fw.txt", 0777)
	fileInfo, err = os.Stat("./fw.txt")
	fileMode = fileInfo.Mode()
	fmt.Println(fileMode)
}
