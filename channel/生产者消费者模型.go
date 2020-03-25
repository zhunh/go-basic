package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int)  {
	for i:=0; i<10;i++ {
		fmt.Println("生产：", i*i)
		out <- i*i
	}
	close(out)
}

func consumer(in <- chan int)  {
	for num := range in {
		fmt.Println("消费者拿到：",num)
		time.Sleep(time.Second)
	}
}

func main()  {
	ch := make(chan int, 6)
	go producer(ch)		// 子go程 生产者
	consumer(ch)		// 主go程 消费
}
