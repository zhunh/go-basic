package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//创建、监听socket
	listenner, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err) //log.Fatal()会产生panic
	}

	defer listenner.Close()

	conn, err := listenner.Accept() //阻塞等待客户端连接
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close() //此函数结束时，关闭连接套接字

	//conn.RemoteAddr().String()：连接客服端的网络地址
	ipAddr := conn.RemoteAddr().String()
	fmt.Println(ipAddr, "连接成功")

	buf := make([]byte, 4096) //缓冲区，用于接收客户端发送的数据

	//阻塞等待用户发送的数据
	n, err := conn.Read(buf) //n代码接收数据的长度
	if err != nil {
		fmt.Println(err)
		return
	}

	//切片截取，只截取有效数据
	result := buf[:n]
	fmt.Printf("接收到数据来自[%s]==>:\n%s\n", ipAddr, string(result))
}