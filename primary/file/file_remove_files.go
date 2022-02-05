package main

import (
	"fmt"
	"os"
)

func main() {
	//创建一个名为“fw_remove”的目录，perm权限为0777
	err := os.Mkdir("fw_remove", 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("创建 目录:fw_remove")
	file1, err := os.Create("./fw_remove/fw_remove_test1.txt")
	if err != nil {
		fmt.Println(err)
	}
	file1.Close()
	fmt.Println("创建 文件:test_remove1.txt")
	file2, err := os.Create("./fw_remove/fw_remove_test2.txt")
	if err != nil {
		fmt.Println(err)
	}
	file2.Close()
	fmt.Println("创建 文件:test_remove2.txt")
	err = os.Remove("./fw_remove/fw_remove_test1.txt")
	if err != nil {
		fmt.Printf("removed ./fw_remove/fw_remove_test1.txt err : %v\n", err)
	}
	fmt.Println("removed 文件:./fw_remove/fw_remove_test1.txt")
	err = os.RemoveAll("./fw_remove")
	if err != nil {
		fmt.Printf("remove all ./fw_remove err : %v\n", err)
	}
	fmt.Println("removed all files:./fw_remove")
}
