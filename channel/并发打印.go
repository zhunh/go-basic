package main

import "fmt"

func printer2(c chan int){
	//开始无限循环等待数据
	for{
		//从channel中获取一个数据
		data := <-c
		//将0视为数据结束
		if data == 0{
			break
		}
		//打印数据
		fmt.Println(data)
	}
	//通知main已经结束循环
	c <- 0
}

func main() {
	c := make(chan int)
	//并发执行printer2,传入channel
	go printer2(c)

	for i:= 1;i<=10;i++{
		//将数据通过channel投送给printer2
		c <- i
		//fmt.Println("已发送",i)
	}
	//通知并发的printer2结束循环（没数据啦！）
	fmt.Println("main循环结束")
	c <- 0
	//等待printer2结束
	<- c
}
