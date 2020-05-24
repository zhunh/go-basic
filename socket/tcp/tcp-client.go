package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)
//tcp客户端，可以多开，服务端会分开处理每个连接
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8010")
	if err != nil{
		fmt.Println("err: ", err)
	}

	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)//os.Stdin 代表标准输入
	for{
		//从终端读取一行数据
		line, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(line, "\r\n")
		if strings.ToUpper(inputInfo) == "Q"{ //输入q退出
			return
		}
		//发送数据
		n, err := conn.Write([]byte(inputInfo)) //发送数据
		if err != nil{
			fmt.Println("发送数据失败, conn.Write err=: ",err)
			//return
		}
		fmt.Printf("客户端发送了 %d 字节的数据。\n", n)

		//buf:= [512]byte{}
		//n, err = conn.Read(buf[:])
		//if err != nil{
		//	fmt.Println("recv failed, err: ", err)
		//	return
		//}
		//
		//fmt.Println(string(buf[:n]))
	}
}