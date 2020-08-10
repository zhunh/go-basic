package main

import (
	"chatroom/client/process"
	"fmt"
	"os"
)

var userId int
var userPwd string
var userName string

func main() {
	//接收用户的选择
	var key int
	for {
		fmt.Println("-----------欢迎登陆多人聊天系统-----------")
		fmt.Println("\t\t\t 1. 登陆聊天室")
		fmt.Println("\t\t\t 2. 注册")
		fmt.Println("\t\t\t 3. 退出")
		fmt.Println("---------------请输入1-3---------------")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆")
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n",&userId)
			fmt.Println("请输入密码:")
			fmt.Scanf("%s\n",&userPwd)
			//完成登陆
			//1.创建userProcess实例
			up := &process.UserProcess{}
			up.Login(userId,userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id：")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码：")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名字(nickname)：")
			fmt.Scanf("%s\n", &userName)
			up := &process.UserProcess{}
			//调用 UserProcess, 完成注册
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出")

			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入。")
		}
	}
//	//根据用户输入，显示新的提示信息
//	if key == 1{
//		//创建
//		up := &process.UserProcess{
//
//		}
//		//先把登陆的函数，写到另外一个文件
//		up.Login(userId,userPwd)
//		//if err != nil{
//		//	fmt.Println("登陆失败, err=", err)
//		//}else {
//		//	fmt.Println("登陆成功")
//		//}
//	}else if key == 2{
//		fmt.Println("进行用户注册")
//	}
}
