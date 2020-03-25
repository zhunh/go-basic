package gorotine

import (
	"fmt"
	"sync"
	"time"
)

//协程同步
//系统工具 sync.waitgroup
//Add(delta int) 添加协程记录
//Done() 移除协程记录
//Wait()同步等待所有记录的协程全部结束

var WG = sync.WaitGroup{}

func Read() {
	for i := 0; i < 3; i++ {
		WG.Add(1)
	}
}

func Write() {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println("Done -> ",i)
		WG.Done()
	}
}