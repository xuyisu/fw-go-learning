package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// 判断在 b 中能否找到正则表达式 pattern 所匹配的子串
	// pattern：要查找的正则表达式
	// b：要在其中进行查找的 []byte
	// matched：返回是否找到匹配项
	// err：返回查找过程中遇到的任何错误
	match, _ := regexp.Match("He.* ", []byte("Hello World!"))
	fmt.Println("Match:", match)
	// 判断在 r 中能否找到正则表达式 pattern 所匹配的子串
	// pattern：要查找的正则表达式
	// r：要在其中进行查找的 RuneReader 接口
	// matched：返回是否找到匹配项
	// err：返回查找过程中遇到的任何错误
	r := bytes.NewReader([]byte("Hello World!"))
	reader, _ := regexp.MatchReader("He.* ", r)
	fmt.Println("MatchReader:", reader)

	// 判断在 s 中能否找到正则表达式 pattern 所匹配的子串
	// pattern：要查找的正则表达式
	// r：要在其中进行查找的字符串
	// matched：返回是否找到匹配项
	// err：返回查找过程中遇到的任何错误
	matchString, _ := regexp.MatchString("He.* ", "Hello World!")
	fmt.Println("Match:", matchString)

	// Compile 用来解析正则表达式 expr 是否合法，如果合法，则返回一个 Regexp 对象
	// FindString 在 s 中查找 re 中编译好的正则表达式，并返回第一个匹配的内容
	compile, err := regexp.Compile(`\w+`)
	fmt.Printf("Compile:")
	fmt.Printf("%q,%v\n", compile.FindString("Hello World!"), err)

	// MustCompile 的作用和 Compile 一样
	// 不同的是，当正则表达式 str 不合法时，MustCompile 会抛出异常
	mustCompile := regexp.MustCompile(`\w+`)
	fmt.Println(mustCompile.FindString("Hello World!"))

	// 在 b 中查找 re 中编译好的正则表达式，并返回第一个匹配的内容
	find := regexp.MustCompile(`\w+`)
	fmt.Printf("%q", find.Find([]byte("Hello World!")))
}
