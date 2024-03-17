package common

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func WritePkg(conn net.Conn, data []byte) (err error) {
	// 先发送一个长度
	pkgLen := uint32(len(data))
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, pkgLen)
	n, err := conn.Write(buf)
	if n != 4 || err != nil {
		fmt.Println("conn.Write fail, err =", err)
		return
	}
	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write fail, err =", err)
		return
	}
	return
}

func ReadPkg(conn net.Conn) (msg Message, err error) {
	buf := make([]byte, 1024)

	// 读取数据包封装成一个函数readPkg
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.Read fail, err =", err)
		return
	}
	// 读取数据包长度，转成uint32
	lenPkg := binary.BigEndian.Uint32(buf[0:4])

	// 根据lenPkg去读取相应的长度，解释下：从conn中读取[:lenPkg]的字节丢在buf中
	n, err := conn.Read(buf[:lenPkg])
	if n != int(lenPkg) || err != nil {
		fmt.Println("conn.Read fail, err =", err)
		return
	}
	// 反序列化成Message类型，下面这个msg前面要取地址
	err = json.Unmarshal(buf[:lenPkg], &msg)
	if err != nil {
		fmt.Println("json.Unmarshal fail, err =", err)
	}
	return
}
