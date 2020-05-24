package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	//循环读取 网页数据，
	buf := make([]byte, 4096)
	for{
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取网页完成")
			break
		}
		if err2 != nil && err2 != io.EOF{
			err = err2
			return
		}
		//累加每一次读到的 buf 数据
		result += string(buf[:n])
	}
	return result, err
}

func SpiderPage(i int, page chan int)  {
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn="+ strconv.Itoa((i-1)*50)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err:", err)
		return
	}
	//fmt.Println(result1)
	//将读取到的网页数据保存成网页
	f, err := os.Create("第"+strconv.Itoa(i)+"页"+".html")
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	f.WriteString(result)
	f.Close()	//保存好一个文件关闭一个文件

	page <- i
}

func working(start, end int) {
	fmt.Println("正在爬取第%d页到第%d页...", start,end)

	page := make(chan int)

	//循环抓取每一页
	for i:=start;i<=end ; i++ {
		go SpiderPage(i, page)
	}
	for i:=start;i<=end ; i++ {
		fmt.Println("第%d个页面爬取完成", <-page)
	}
}

func main() {
	//指定爬取的起始页
	var start, end int
	fmt.Println("请输入爬去的起始页（>=1）:")
	fmt.Scan(&start)
	fmt.Println("请输入爬去的终止页（>=1）:")
	fmt.Scan(&end)

	working(start, end)
}
