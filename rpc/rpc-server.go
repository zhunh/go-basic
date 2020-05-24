package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

func httpHandle(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w,"its rpc test.")
}
type RpcObj int
//函数关键字 （对象） 函数名 （对端发送过来的内容，返回给对端的内容）错误
func (this *RpcObj) GetInfo(argType int, replyType *int) error{
	fmt.Println("打印对端发送过来的内容为：", argType)
	//修改内容值
	*replyType = argType + 1
	return  nil
}
func main(){
	//页面请求
	http.HandleFunc("/rpctest", httpHandle)
	//实例化对象
	pd := new(RpcObj)
	//服务端注册一个对象
	rpc.Register(pd)

	rpc.HandleHTTP()
	//获取连接
	ln, err := net.Listen("tcp", ":10086")
	if err != nil{
		fmt.Println("网络错误")
	}
	http.Serve(ln, nil)
}
