package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Mkdir("fw_dir", 0777)
	if err != nil {
		fmt.Println(err)
	}
	oldName := "fw_dir"
	newName := "fw_tmp"
	//将fw_dir重命名为fw_tmp
	err = os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}
