package main

import (
	"fmt"
	"reflect"
)

//反射
func main() {
	var x float64 = 3.4
	//fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))
}
