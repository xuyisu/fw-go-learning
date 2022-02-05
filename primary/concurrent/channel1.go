package main

func main() {
	// 创建一个字符串通道
	ch := make(chan string)
	// 尝试将sleep通过通道发送
	ch <- "sleep"

	//会返回 all goroutines are asleep - deadlock!
}
