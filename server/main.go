package main

import (
	"fmt"
	"net"
)

// 处理与客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()
	// 循环读取客户端发送的消息
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf[:4])
		if err != nil {
			fmt.Println("conn.Read fail, err =", err)
			return
		}
		fmt.Println("读到的buf信息是:", buf[:4])
	}

}

func main() {
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen fail, err =", err)
		return
	}
	for {
		fmt.Println("等待客户端连接。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("fail listen.Accept, err =", err)
		}
		// 一旦链接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}
}
