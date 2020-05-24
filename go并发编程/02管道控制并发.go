package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

/*
	启动100个协程计算1~100的平方根
	要求同时只有5个协程在执行
*/

func getSqrt(name string, n int, ch chan string)  {
	//先注册到通道，能写入就执行，不能就阻塞到能写入为止
	ch<- name
	ret := math.Sqrt(float64(n))
	time.Sleep(time.Second)
	fmt.Printf("%d的平方根是%.2f\n", n, ret)

	//任务执行完毕，从信号量控制管道注销自己，以便为其他并发任务腾出空间
	<-ch
}

func main() {
	//创建一个缓存为5的通道
	chSem := make(chan string, 5)
	for i:=1; i<100; i++{
		go getSqrt("协程"+strconv.Itoa(i), i, chSem)
	}
	time.Sleep(time.Second*10)
}
/*
运行结果可以看见，主线程开辟100个协程，但同时最多有5个协程在工作
*/