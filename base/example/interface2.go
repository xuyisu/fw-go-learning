package main

import "fmt"

//接口多态
type Message interface {
	sending()
}

// 定义user及user.notify方法
type UserMsg struct {
	name  string
	phone string
}

func (u *UserMsg) sending() {
	fmt.Printf("Sending user phone to %s<%s>\n", u.name, u.phone)
}

// 定义admin及admin.message方法
type adminMsg struct {
	name  string
	phone string
}

func (a *adminMsg) sending() {
	fmt.Printf("Sending admin phone to %s<%s>\n", a.name, a.phone)
}

func main() {
	// 创建一个user值并传给sendMessage
	test1 := UserMsg{"test1", "16600002222"}
	sendMessage(&test1)

	// 创建一个admin值并传给sendMessage
	test2 := adminMsg{"test2", "16600002233"}
	sendMessage(&test2)
}

// sendMessage接受一个实现了message接口的值 并发送通知
func sendMessage(n Message) {
	n.sending()
}
