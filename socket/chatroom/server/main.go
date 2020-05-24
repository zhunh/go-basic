package main

import "fmt"
import "server/mi"

var userId int
var userPwd string

func main() {
	//接收用户的选择
	var key int
	//判断是否继续显示菜单
	var loop = true
	for{
		fmt.Println("-----------欢迎登陆多人聊天系统-----------")
		fmt.Println("\t\t\t 1. 登陆聊天室")
		fmt.Println("\t\t\t 2. 注册")
		fmt.Println("\t\t\t 3. 退出")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆")
			loop = false
		case 2:
			fmt.Println("注册")
			loop = false
		case 3:
			fmt.Println("退出")
			loop = false
		default:
			fmt.Println("你的输入有误，请重新输入。")
		}
	}
	//根据用户输入，显示新的提示信息
	if key == 1{
		fmt.Println("请输入用户id")
		fmt.Scanf("%d\n",&userId)
		fmt.Println("请输入密码")
		fmt.Scanf("%d\n",&userPwd)
		//先把登陆的函数，写到另外一个文件
		err := mi.login(userId,userPwd)
		if err != nil{
			fmt.Println("登陆失败")
		}else {
			fmt.Println("登陆失败")
		}
	}else if key == 2{
		fmt.Println("进行用户注册")
	}
}
