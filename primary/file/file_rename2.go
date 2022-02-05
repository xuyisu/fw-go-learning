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
	//创建一个名为“rename_test”的目录，perm权限为0777
	err = os.Mkdir("rename_test", 0777)
	//将fw_rename.txt移动到rename_test目录，并将名字重命名为fw_rename_new.txt
	err = os.Rename("./fw_rename.txt", "./rename_test/fw_rename_new.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
}
