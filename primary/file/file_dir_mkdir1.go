package main

import (
	"fmt"
	"os"
)

func main() {
	//根据path创建目录，例如fw_dir
	err := os.Mkdir("fw_dir", 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("mkdir success")
}
