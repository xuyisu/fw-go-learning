package main

import "fmt"

//接口组合
//邮件
type MessageEmail interface {
	sendingEmail()
}

//微信
type MessageWeChat interface {
	sendingWeChat()
}

//邮件
type Messages interface {
	MessageEmail
	MessageWeChat
}

// 定义user及user.notify方法
type UserCompose struct {
	name  string
	phone string
}

func (u *UserCompose) sendingEmail() {
	fmt.Printf("sendingEmail user phone to %s<%s>\n", u.name, u.phone)
}

func (u *UserCompose) sendingWeChat() {
	fmt.Printf("sendingWeChat user phone to %s<%s>\n", u.name, u.phone)
}

// 定义admin及admin.message方法
type adminCompose struct {
	name  string
	phone string
}

func (a *adminCompose) sendingEmail() {
	fmt.Printf("sendingEmail admin phone to %s<%s>\n", a.name, a.phone)
}
func (a *adminCompose) sendingWeChat() {
	fmt.Printf("sendingWeChat admin phone to %s<%s>\n", a.name, a.phone)
}

func main() {
	// 创建一个user值并传给sendMessage
	test1 := UserCompose{"test1", "15500005555"}
	sendMessageCompose(&test1)

	// 创建一个admin值并传给sendMessage
	test2 := adminCompose{"test2", "15511115555"}
	sendMessageCompose(&test2)
}

// sendMessage接受一个实现了message接口的值 并发送通知
func sendMessageCompose(n Messages) {
	n.sendingEmail()
	n.sendingWeChat()
}
