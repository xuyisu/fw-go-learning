package main

import (
	"bytes"
	"fmt"
)

//Buffer拼接数据(线程不安全)
func main() {
	var buffer bytes.Buffer
	for i := 0; i < 10; i++ {
		buffer.WriteString("a")
	}
	fmt.Println(buffer.String())
}
