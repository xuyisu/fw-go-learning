package main

import (
	"fmt"
	"fw-go-web/base/entity"
)

func main() {
	s := new(entity.Student)
	//s.name="test"
	s.Age = 25
	s.SetName("666")
	fmt.Println(s.Age)
	fmt.Println(s.GetName())
}
