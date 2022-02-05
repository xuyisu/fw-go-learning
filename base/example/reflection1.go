package main

import (
	"fmt"
	"reflect"
)

//反射
func main() {
	var name string = "学习go"

	v := reflect.ValueOf(name)
	fmt.Println("可写性为:", v.CanSet())
}
