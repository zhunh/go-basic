package main

import (
	"fmt"
	"math/rand"
	"time"
)

func readGo05(in <-chan int,idx int)  {
	for{
		num := <- in			//从channel读出数据
		fmt.Printf("--%dth 读go程， 读出：%d\n", idx,num)
		time.Sleep(time.Second)
	}
}

func writeGo05(out chan<- int, idx int)  {
	for{
		//生成随机数
		num := rand.Intn(100)
		out <- num				//写入channel
		fmt.Printf("%dth 写go程，写入：%d\n", idx,num)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	//播种随机数种子
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go readGo05(ch,i+1)
	}
	for i := 0; i < 5; i++ {
		go writeGo05(ch,i+1)
	}

	for{
		;
	}
}
