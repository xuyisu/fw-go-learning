package main

import (
	"fmt"
	"reflect"
)

//利用反射修改数据
func main() {
	var name string = "张三"
	fmt.Println("真实 name 的原始值为：", name)

	v1 := reflect.ValueOf(&name)
	v2 := v1.Elem()

	v2.SetString("刘备")
	fmt.Println("通过反射对象进行更新后，真实 name 变为：", name)
}
