package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建服务器UDP地址结构。指定 IP + port
	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	// 监听客户端连接
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		fmt.Println("net.ListenUDP err:", err)
		return
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, raddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("conn.ReadFromUDP err:", err)
			return
		}
		fmt.Printf("接收到客户端信息 [%s]: %s", raddr, string(buf[:n]))

		conn.WriteToUDP([]byte("服务端已收到"), raddr) // 简单回写数据给客户端
	}
}
