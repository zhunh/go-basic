package main

import (
	"fmt"
	"sync"
	"time"
)

func main01() {
	var wg sync.WaitGroup
	fmt.Println(wg)

	//向等待组中添加一个协程(注册)
	wg.Add(1)

	//从等待组中减掉一个协程(注销)
	wg.Done()

	//阻塞等待组中的协程数归零
	wg.Wait()
}
/*
分别使用Ticker和Timer创建耗时协程A,B,C
将A，B，C三个协程加入等待组
A,B,C结束时从等待组注销
主协程阻塞至等待组内协程数归零
*/
func main()  {
	//生命等待组
	var wg sync.WaitGroup
	//每开辟一个协程就向等待组中+1
	wg.Add(1)
	go func() {
		fmt.Println("协程1开始工作")
		time.Sleep(time.Second*4)
		fmt.Println("协程1结束")
		//协程结束，从等待组中-1
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("协程2开始工作")
		tk := time.NewTimer(2*time.Second)
		<-tk.C
		fmt.Println("协程2结束")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("协程3开始工作")
		tk := time.NewTicker(1*time.Second)
		for i:=0;i<5;i++{
			<-tk.C
		}
		fmt.Println("协程3结束")
		wg.Done()
	}()

	//阻塞等到wg中的协程归零
	wg.Wait()
	fmt.Println("main over")
}