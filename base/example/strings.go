package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "你好,世界!!!!,hello,world"
	// Count 计算字符串 sep 在 s 中的非重叠个数
	count := strings.Count(s, "!")
	fmt.Println("Count统计结果:", count)
	//Contains 判断字符串 s 中是否包含子串 substr  如果 substr 为空，则返回 true
	contains := strings.Contains(s, "你好")
	contains2 := strings.Contains(s, "")
	fmt.Println("Contains是否包含你好:", contains)
	fmt.Println("Contains空包含:", contains2)
	// ContainsAny 判断字符串 s 中是否包含 chars 中的任何一个字符  如果 chars 为空，则返回 false
	containsAny := strings.ContainsAny(s, ",")
	fmt.Println("ContainsAny是否包含,:", containsAny)
	//ContainsRune 判断字符串 s 中是否包含字符 r
	containsRune := strings.ContainsRune(s, '你')
	fmt.Println("ContainsRune是否包含,:", containsRune)
	//Index 返回子串 sep 在字符串 s 中第一次出现的位置,如果找不到，则返回 -1，如果 sep 为空，则返回 0。
	index := strings.Index(s, "w")
	fmt.Println("Index首次索引,:", index)
	//LastIndex 返回子串 sep 在字符串 s 中最后一次出现的位置,如果找不到，则返回 -1，如果 sep 为空，则返回 0。
	lastIndex := strings.LastIndex(s, "h")
	fmt.Println("Index最后一次索引:", lastIndex)
	// IndexAny 返回字符串 chars 中的任何一个字符在字符串 s 中第一次出现的位置 如果找不到，则返回 -1，如果 chars 为空，则返回 -1
	indexAny := strings.IndexAny(s, "hel")
	fmt.Println("IndexAny索引:", indexAny)
	// LastIndexAny 返回字符串 chars 中的任何一个字符在字符串 s 中最后一次出现的位置
	lastIndexAny := strings.LastIndexAny(s, "wor")
	fmt.Println("LastIndexAny索引:", lastIndexAny)
	// SplitN 以 sep 为分隔符，将 s 切分成多个子串，结果中不包含 sep 本身  参数 n 表示最多切分出几个子串，超出的部分将不再切分
	splitN := strings.SplitN(s, ",", 5)
	fmt.Println("SplitN切的数据:", splitN)
	// SplitAfterN 以 sep 为分隔符，将 s 切分成多个子串，结果中包含 sep 本身,参数 n 表示最多切分出几个子串，超出的部分将不再切分。
	splitAfterN := strings.SplitAfterN(s, ",", 5)
	fmt.Println("SplitAfterN切的数据:", splitAfterN)
	// Split 以 sep 为分隔符，将 s 切分成多个子切片，结果中不包含 sep 本身
	split := strings.Split(s, ",")
	fmt.Println("Split切的数据:", split)
	// SplitAfter 以 sep 为分隔符，将 s 切分成多个子切片，结果中包含 sep 本身
	splitAfter := strings.SplitAfter(s, ",")
	fmt.Println("SplitAfter切的数据:", splitAfter)
	// Join 将 a 中的子串连接成一个单独的字符串，子串之间用 sep 分隔
	ss := []string{"one", "two", "three"}
	join := strings.Join(ss, "|")
	fmt.Println(join)

}
