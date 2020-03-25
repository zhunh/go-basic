package gorotine

import (
	"fmt"
	"time"
)

var chanInt chan int = make(chan int, 10)

var chanBool chan bool = make(chan bool, 10)
//发送数据到chan
func Send()  {
	time.Sleep(time.Second * 1)
	chanInt <- 1
	time.Sleep(time.Second * 1)
	chanInt <- 2
	time.Sleep(time.Second * 1)
	chanInt <- 3
	time.Sleep(time.Second * 2)
	chanBool <- true
}
//接受chan中的数据
func Recv()  {
	//num := <- chanInt
	//fmt.Println("num:",num)
	//num = <- chanInt
	//fmt.Println("num:",num)
	//num = <- chanInt
	//fmt.Println("num:",num)
	//select将来自不同通道的数据分类
	for{
		select {
		case num := <- chanInt:
			fmt.Println("data from chanInt:",num)
		case value := <- chanBool:
			fmt.Println("data from chanBool:",value)
		}
	}
}