package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn)  {
	defer conn.Close()
	for{
		reader := bufio.NewReader(conn)

		var buf [128]byte
		n, err := reader.Read(buf[:]) //读取数据
		if err != nil{
			fmt.Println("read from client failed, err:", err)
			break
		}

		recvStr := string(buf[:n])
		fmt.Println("recv data from client: ",recvStr)

	}
}

func process02(conn net.Conn)  {
	defer conn.Close()

	for{
		buf := make([]byte, 1024)
		//服务器等待客户端发送数据
		//fmt.Printf("服务器等待客户端 %s 发送数据.\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil{
			fmt.Println("服务器的Reader err=", err)
			return
		}
		//显示客户端发送的数据
		fmt.Println(string(buf[:n]))
	}
}

func main()  {
	listener, err := net.Listen("tcp", "127.0.0.1:8010")
	if err != nil{
		fmt.Println("net.Listen error.", err)
		return
	}

	for{
		fmt.Println("等待客户端连接...")
		conn, err := listener.Accept() //建立连接
		if err != nil{
			fmt.Println("accept failed, err:", err)
			continue
		}
		fmt.Printf("New conn suc con=%v. client ip = %v\n",conn, conn.RemoteAddr())
		go process02(conn)
	}
}
