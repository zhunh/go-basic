package main

import (
	"fmt"
	"net"
)

func main() {
	//创建套接字与服务器通信
	conn, err := net.Dial("tcp", "127.0.0.1:8002")
	if err !=nil{
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	//主动发数据给服务器
	conn.Write([]byte("are u ready?"))
	//接收服务端发送的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err !=nil{
		fmt.Println("conn.Read err:", err)
		return
	}
	fmt.Println("服务器回发：",string(buf[:n]))

}
