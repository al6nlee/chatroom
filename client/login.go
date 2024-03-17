package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	common "github.com/al6nlee/chatroom/common/message"
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
	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write fail, err =", err)
	}
	msg, err = common.ReadPkg(conn)

	if err != nil {
		fmt.Println("common.ReadPkg fail, err =", err)
	}
	var resMsgLogin common.LoginResultMessage
	err = json.Unmarshal([]byte(msg.Data), &resMsgLogin)
	if err != nil {
		fmt.Println("json.Unmarshal fail, err =", err)
		return
	}
	if resMsgLogin.Code == 200 {
		fmt.Println("登陆成功啦")
	} else {
		fmt.Println("登录失败, err =", resMsgLogin.Error)
	}
	// 这里还需要服务器端返回的消息
	return
}
