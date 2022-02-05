package main

import (
	"fmt"
	"os"
)

func main() {
	//新建文件
	fileOut, err := os.Create("./fw_write.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileOut.Close()
	for i := 0; i < 5; i++ {
		outStr := fmt.Sprintf("%s:%d\r\n", "Hello Go", i)
		// 写入文件
		fileOut.WriteString(outStr)            //string信息
		fileOut.Write([]byte("i love go\r\n")) //byte类型
		fileOut.WriteAt([]byte("插入一句话"), 5)    //从指定位置写
	}
}
