package main

import (
	"fmt"
	"net"
)

func main() {
	add := net.UDPAddr{
		IP: net.IPv4(0,0,0,0),
		Port:8020,
	}
	listener, err := net.ListenUDP("udp", &add)
	if err != nil{
		fmt.Println("listen failed, err: ", err)
		return
	}

	defer listener.Close()

	for{
		var data [1024]byte
		n, addr, err := listener.ReadFromUDP(data[:]) //接收数据
		if err != nil{
			fmt.Println("read udp from failed, err: ", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)

		_, err = listener.WriteToUDP(data[:n], addr) //发送数据
		if err != nil{
			fmt.Println("write to udp failed, err: ", err)
			continue
		}
	}
}
