package main

import (
	"fmt"
	"os"
)

func main() {
	// 创建文件
	//文件的创建，Create会根据传入的文件名创建文件，默认权限是0666
	fp, err := os.Create("./test.txt") // 如果文件已存在，会将文件清空。
	if err != nil {
		fmt.Println("文件创建失败。")
		//如果文件创建失败，一般是：路径不存在 、权限不足、打开文件数量超过上限、磁盘空间不足等原因
		return
	}

	// defer延迟调用
	defer fp.Close() //关闭文件，释放资源。
}
