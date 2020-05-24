package main

import (
	"google.golang.org/grpc"
	"fmt"
	pd "going/grpc/myproto"
	"context"
)

func main() {
	//客户端连接服务器
	conn ,err :=grpc.Dial("127.0.0.1:10086",grpc.WithInsecure())
	if err!=nil {
		fmt.Println("网络异常",err)
	}
	//网络延迟关闭
	defer  conn.Close()


	//获得grpc句柄
	c:=pd.NewHelloserverClient(conn)

	//通过句柄调用函数
	re ,err :=c.Sayhello(context.Background(),&pd.HelloReq{Name:"Kristin"})
	if err!=nil {
		fmt.Println("sayhello 服务调用失败")
	}
	fmt.Println("调用sayhello的返回",re.Msg)



	re1 ,err :=c.Sayname(context.Background(),&pd.NameReq{Name:"Rona"})
	if err !=nil{
		fmt.Println("say name 调用失败")
	}
	fmt.Println("调用Sayname的返回",re1.Msg)

}