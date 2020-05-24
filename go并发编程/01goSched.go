package main

import (
	"fmt"
	"runtime"
	"time"
)

//出让协程资源

func print(no int) {
	//fmt.Println("子协程初始化完毕..")
	if no == 5{
		//出让当前协程的资源
		runtime.Gosched()
	}
	for j:= 1; j<=10; j++{
		fmt.Printf("协程%d: %d\n", no, j)
	}
}
func main1() {
	for i := 1; i<=9; i++{
		go print(i)
	}
	time.Sleep(time.Second*1)
}

func main()  {
	//设置最大可用CPU逻辑核心数，并返回之前的 cpu核心数设置
	//在当前协程生效
	//runtime.GOMAXPROCS(3)
	fmt.Println("当前可用cpu核心数：", runtime.GOMAXPROCS(3))//4
	fmt.Println("当前可用cpu核心数：", runtime.GOMAXPROCS(1))//3
	//获得逻辑cpu 核心数
	cpu_num := runtime.NumCPU()
	fmt.Println("cpu 个数：", cpu_num)
}