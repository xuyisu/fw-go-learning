package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("fw.txt")
	if err != nil {
		fmt.Println(err)
	}
	// 关闭文件句柄
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		// 读到一个换行符就结束
		line, err := reader.ReadString('\n')
		// io.EOF表示文件的末尾
		if err == io.EOF {
			break
		}
		// 输出内容
		fmt.Print(line)
	}
}
