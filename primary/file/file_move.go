package main

import (
	"fmt"
	"os"
)

//移动文件，可以修改名称
func main() {
	err := os.Rename("./file1", "/tmp/file1")
	if err != nil {
		fmt.Println(err)
		return
	}
}
