package main

import (
	"fmt"
	"os"
)

func main() {
	//根据path创建多级子目录，例如fw_dir/1/2/3
	err := os.MkdirAll("fw_dir/1/2/3", 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("mkdir success")
}
