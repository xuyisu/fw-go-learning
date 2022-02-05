package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	printGo()
	printStr()
}

func printGo() {
	str := "\"go\" 编程学习,you can do it"
	str2 := `"go 编程学习",
you can do it`

	fmt.Println(str)
	fmt.Println(str2)
}

func printStr() {
	str := "you can do it"
	fmt.Println(str[1])                      //获取索引为1的原始字节
	fmt.Println(str[1:5])                    //获取索引1到5的原始字节
	fmt.Println(str[5:])                     //获取索引5到结尾的原始字节
	fmt.Println(string(str[1]))              //获取索引为1的字符串
	fmt.Println([]rune(str))                 //获取每个字节转换的码点值
	fmt.Println(len(str))                    //获取字符串的字节数
	fmt.Println(utf8.RuneCountInString(str)) //获取字符串的字符数
}
