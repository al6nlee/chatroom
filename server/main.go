package main

import (
	"encoding/json"
	"errors"
	"fmt"
	common "github.com/al6nlee/chatroom/common/message"
	"io"
	"net"
)

func serverProcessLogin(conn net.Conn, msg *common.Message) (err error) {
	// 从msg中获取data并反序列化
	var loginMsg common.LoginMessage
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("json.Unmarshal fail, err =", err)
		return
	}
	// 声明一个resMsg, 用于返回客户端
	var resMsg common.Message
	resMsg.Type = common.LoginResultMessageType
	var resLoginMsg common.LoginResultMessage

	// 判断用户id、pwd是否合法
	if loginMsg.UserId == 100 && loginMsg.UserPwd == "1234567" {
		// 合法
		resLoginMsg.Code = common.SucessCode
		fmt.Println("用户、密码合法")
	} else {
		// 不合法
		resLoginMsg.Code = common.FailCode
		resLoginMsg.Error = "用户、密码不合法"
	}
	// 序列化数据
	data, err := json.Marshal(resLoginMsg)
	if err != nil {
		fmt.Println("json.Marshal fail, err =", err)
		return
	}
	resMsg.Data = string(data)
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marshal fail, err =", err)
		return
	}
	// 发送给客户端
	err = common.WritePkg(conn, data)
	if err != nil {
		fmt.Println("writePkg fail, err =", err)
		return
	}
	return
}

// 根据客户端发送的消息类型，去做不同的处理
func serverProcessMsg(conn net.Conn, msg *common.Message) (err error) {
	switch msg.Type {
	case common.LoginMessageType:
		// 处理登录逻辑
		err = serverProcessLogin(conn, msg)
	default:
		fmt.Println("消息类型不存在，无法处理")
		errors.New("消息类型不存在，无法处理")
	}
	return
}

// 处理与客户端的通讯
func process(conn net.Conn) {
	defer conn.Close()
	// 循环读取客户端发送的消息
	for {
		msg, err := common.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端已退出，服务端也退出")
				return
			}
			fmt.Println("readPkg fail, err =", err)
		}
		serverProcessMsg(conn, &msg)
		fmt.Println(msg)
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
