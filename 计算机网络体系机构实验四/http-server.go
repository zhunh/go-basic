package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

func httpHandle(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w,"its rpc test.")
}

func main(){
	//页面请求
	http.HandleFunc("/rpctest", httpHandle)

	//服务端注册一个对象
	ln, err := net.Listen("tcp", ":10086")
	if err != nil{
		fmt.Println("网络错误")
	}
	http.Serve(ln, nil)
}
