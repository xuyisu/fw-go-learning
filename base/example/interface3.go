package main

import "fmt"

//类型转换
type MessageType interface {
	sending()
}

// 定义user及user.notify方法
type UserType struct {
	name  string
	phone string
}

func (u UserType) sending() {
	fmt.Printf("Sending user phone to %s<%s>\n", u.name, u.phone)
}

// 定义admin及admin.message方法
type AdminType struct {
	name  string
	phone string
}

func (a AdminType) sending() {
	fmt.Printf("Sending admin phone to %s<%s>\n", a.name, a.phone)
}

func main() {
	var s MessageType
	// 创建一个user值并传给sendMessage
	s = UserType{"test1", "16600002222"}
	sendMessageTypeBySwitch(s)
	sendMessageTypeByIf(s)

	// 创建一个admin值并传给sendMessage
	s = AdminType{"test2", "16600002233"}
	sendMessageTypeBySwitch(s)
	sendMessageTypeByIf(s)
}

// switch方式 的类型转换
func sendMessageTypeBySwitch(n MessageType) {
	switch instance := n.(type) {
	case UserType:
		fmt.Println("switch user", instance.name)
		n.sending()
	case AdminType:
		fmt.Println("switch admin", instance.name)
		n.sending()
	}
}

// if方式 的类型转换
func sendMessageTypeByIf(n MessageType) {
	if instance, ok := n.(UserType); ok {
		fmt.Println("if user", instance.name)
		n.sending()
	} else if instance, ok := n.(AdminType); ok {
		fmt.Println("if admin", instance.name)
		n.sending()
	}
}
