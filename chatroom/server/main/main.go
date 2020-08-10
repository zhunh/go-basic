package main

import (
	"chatroom/server/model"
	"chatroom/server/utils"
	"fmt"
	"net"
	"chatroom/server/process"
)

//处理和客户端的通讯
func processor(conn net.Conn)  {
	defer conn.Close()
	//调用总控 process
	processor := &process.Processor{
		Conn:conn,
	}
	err := processor.Processor()
	if err != nil{
		fmt.Println("客户端和服务器通讯的协程错误，err=",err)
		return
	}
}
func init()  {
	//程序启动，初始化连接池
	utils.InitPool()
	InitUserDao()
}
//初始化一个 userDao
func InitUserDao()  {
	//这里的pool 是在redis.go 定义的全局变量
	model.MyUserDao = model.NewUserDao(utils.Pool)
}
func main() {
	//提示信息
	fmt.Println("服务器在8889端口监听....")
	listen, err := net.Listen("tcp","0.0.0.0:8889")
	defer listen.Close()
	if err != nil{
		fmt.Println("net.listen err, err=",err)
		return
	}
	//一旦监听成功，就循环等待客户端连接
	for{
		conn, err := listen.Accept()
		if err != nil{
			fmt.Println("listen.Accept err, err=",err)
		}
		//一旦连接成功，则启动一个协程和客户端保持通讯。
		go processor(conn)
	}
}
