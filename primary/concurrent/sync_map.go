package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("zf", "张飞")
	scene.Store("gy", "关羽")
	scene.Store("lb", "刘备")
	scene.Store("zl", "张良")
	scene.Store("cc", "曹操")
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("zf"))
	// 根据键删除对应的键值对
	scene.Delete("lb")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
