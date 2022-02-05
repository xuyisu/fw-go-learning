package main

import (
	"fmt"
	"os"
)

func main() {
	//先创建多级子目录
	os.MkdirAll("fw_dir/1/2/3", 0777)
	//删除fw_dir目录及其子目录
	err := os.RemoveAll("fw_dir")
	if err != nil {
		fmt.Printf("remove fw_dir err : %v\n", err)
	}
	fmt.Println("remove success")

}
