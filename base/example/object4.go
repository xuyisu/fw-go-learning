package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Man struct {
	Human
	game string
}

type Woman struct {
	Human
	skirt string
}

func (user *Human) excute() {
	fmt.Printf("用户姓名:%s 年龄:%d\n", user.name, user.age)
}
func (man *Man) excute() {
	fmt.Printf("用户姓名:%s 年龄:%d,想法:%s\n", man.name, man.age, man.game)
}
func (woman *Woman) excute() {
	fmt.Printf("用户姓名:%s 年龄:%d,想法:%s\n", woman.name, woman.age, woman.skirt)
}

func main() {
	man := Man{Human{name: "刘备", age: 60}, "我太喜欢玩游戏了"}
	woman := Woman{Human{name: "貂蝉", age: 17}, "我太喜欢买裙子了"}
	man.excute()
	woman.excute()
}
