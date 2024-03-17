package main

import (
	"encoding/binary"
	"encoding/json"
	common "example.com/greetings/common/message"
	"fmt"
	"net"
)

func login(userId int, userPwd string) (err error) {
	// 建立连接
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial fail, err =", err)
		return err
	}
	// 延时关闭
	defer conn.Close()

	// 准备消息
	var msg common.Message
	msg.Type = common.LoginMessageType

	var loginMsg common.LoginMessage
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd

	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("json.Marshal fail, err =", err)
	}
	msg.Data = string(data)

	data, err = json.Marshal(msg)
	// 先发送data的长度，而conn.Write() 接受的是切片
	pkgLen := uint32(len(data))
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, pkgLen)
	n, err := conn.Write(buf)
	if n != 4 || err != nil {
		fmt.Println("conn.Write fail, err =", err)
	}
	fmt.Println("客户端输入的长度=", len(data), "内容为", string(data))
	return nil
}
