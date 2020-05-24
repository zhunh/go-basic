package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//resp, err := http.Get("http://www.baidu.com/")
	resp, err := http.Get("http://127.0.0.1:10086/rpctest") //获取本地服务器页面数据
	if err != nil {
		// handle error
		fmt.Println("访问出错")
	}
	defer resp.Body.Close()
	head := resp.Header
	fmt.Println(head)
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	//写入文件
	data := []byte(string(body))
	if ioutil.WriteFile("test2.html",data,0644) == nil{
		fmt.Println("写入文件成功")
	}
}
