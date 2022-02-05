package main

import (
	"fmt"
	"io"
	"os"
)

//文件复制
func main() {
	//先创建一个名为：fw_copy.txt文件
	file, err := os.Create("./fw_copy.txt")
	if err != nil {
		fmt.Println(err)
	}
	file.WriteString("Hello Go")
	//打开文件fw_copy.txt，获取文件指针
	srcFile, err := os.Open("./fw_copy.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer srcFile.Close()
	//打开文件要复制的新文件名fw_copy2.txt，获取文件指针
	dstFile, err := os.OpenFile("./fw_copy2.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}
	defer dstFile.Close()
	//通过Copy方法
	result, err := io.Copy(dstFile, srcFile)
	if err == nil {
		fmt.Println("复制完成，共复制字节数为: ", result)
	}
}
