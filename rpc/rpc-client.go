package main

import (
	"fmt"
	"net/rpc"
)


func main() {
	//建立连接
	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:10086")
	if err != nil{
		fmt.Println("网络连接失败")
	}
	//定义变量接收远端数据
	var data int
	err = cli.Call("RpcObj.GetInfo",8888,&data)
	if err != nil{
		fmt.Println("远程调用失败")
	}
	fmt.Println("从远端得到的数据为：", data)
}
