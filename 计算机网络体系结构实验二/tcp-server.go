package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp","127.0.0.1:8002")
	if err != nil {
		fmt.Println("net.Listen error.", err)
		return
	}
	defer listener.Close()
	fmt.Println("服务器等待客户端连接请求...")
	//阻塞监听客户端连接请求
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept error.", err)
		return
	}
	defer conn.Close()
	fmt.Println("服务器与客户端连接成功！！！")
	//读取客户端数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read error.", err)
		return
	}
	//处理数据
	fmt.Println("从客户端收到数据：",string(buf[:n]))
	//发动数据给客户端
	conn.Write([]byte("yes ready!"))
}