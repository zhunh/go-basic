package main

import (
	"bytes"
	"fmt"
)

//定义一个函数，参数数量为0~n，类型约束为字符串
func joinString(slist ...string) string {
	//定义一个字节缓冲，快速地连接字符串
	var b bytes.Buffer
	//遍历可变参数列表slist，类型为[]string
	for _, s := range slist{
		//将遍历出的字符串连续写入字节数组
		b.WriteString(s)
	}
	//将连接好的字节数组转换为字符串并输出
	return b.String()
}

func main() {
	//输入3个字符串，将他们连成一个字符串
	fmt.Println(joinString("云村 ","star ","战斗"))
}
