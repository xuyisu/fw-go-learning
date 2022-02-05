package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//先创建多级子目录
	os.MkdirAll("fw_dir/1/2/3", 0777)
	//遍历fw_dir目录及其子目录
	err := filepath.Walk("fw_dir", print)
	if err != nil {
		fmt.Printf("walk fw_dir err : %v\n", err)
	}
	fmt.Println("walk success")

}

func print(path string, info os.FileInfo, err error) error {
	fmt.Println(path)
	return nil
}
