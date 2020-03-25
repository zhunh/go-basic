package main

import (
	"fmt"
	"runtime"
	"time"
)

// 全局定义channel， 用来完成数据同步
var channel = make(chan int)

// 定义一台打印机
func printer(s string)  {
	for _, ch := range s {
		fmt.Printf("%c", ch)			// 屏幕：stdout
		time.Sleep(300 * time.Millisecond)
	}
}

// 定义两个人使用打印机
func person1()  {				// person1 先执行。
	printer("hello")
	channel <- 8	//写入数据到channel
}
func person2()  {				// person2 后执行
	<- channel
	printer("wrold")
}

func main()  {
	go person1()
	//<- channel  //从channel读出数据

	go person2()

	runtime.Goexit()
}
