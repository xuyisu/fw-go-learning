package main

//Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

import "fmt"

func main() {

	//数组的遍历
	numsArr := [5]int{1, 2, 3, 4, 5}
	//遍历数组
	for i, num := range numsArr {
		fmt.Printf("index:%d,value=%d\n", i, num)
	}

	//切片的遍历
	numsSlice := []int{2, 3, 4}
	//遍历数组
	for i, num := range numsSlice {
		fmt.Printf("index:%d,value=%d\n", i, num)
	}
	//range遍历map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//遍历字符串
	for i, c := range "GoLand" {
		fmt.Printf("结果: %d  %c\n", i, c)
	}

	//遍历通道
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}
