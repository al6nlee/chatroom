package main

import (
	"fmt"
	"os"
)

var userId int
var userPwd string

func main() {
	// 定义用户输入指令
	var key int
	// 判断是否继续显示菜单
	var loop bool = true
	for loop {
		fmt.Println("------------------欢迎登录多人聊天系统------------------")
		fmt.Println("\t\t\t1 登录聊天室")
		fmt.Println("\t\t\t2 注册用户")
		fmt.Println("\t\t\t3 退出系统")
		fmt.Println("\t\t\t请选择(1-3)")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
	// 根据用户输入的指令，去往下一个页面
	if key == 1 {
		fmt.Println("请输入用户id")
		fmt.Scanln(&userId)
		fmt.Println("请输入用户密码")
		fmt.Scanln(&userPwd)
		login(userId, userPwd)
	} else if key == 2 {
		fmt.Println("用户注册逻辑")
	}
}
