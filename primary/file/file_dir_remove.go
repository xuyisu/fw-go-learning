package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("fw_dir", 0777)
	err := os.Remove("fw_dir")
	if err != nil {
		fmt.Printf("remove fw_dir err : %v\n", err)
	}
	fmt.Println("remove success")
}
