package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex sync.RWMutex		//锁只有一把，两个属性 r w

func readGo(in <-chan int, idx int)  {
	for{
		rwMutex.RLock()		//以读模式枷锁
		num := <-in
		fmt.Printf("--%dth 读go程， 读出：%d\n", idx,num)
		rwMutex.RUnlock()	//以读模式解锁
	}
}

func writeGo(out chan<- int, idx int)  {
	for{
		//生成随机数
		num := rand.Intn(100)

		rwMutex.Lock()		//以写模式加锁
		out <- num
		fmt.Printf("%dth 写go程，写入：%d\n", idx,num)
		time.Sleep(time.Millisecond * 300)
		rwMutex.Unlock()
	}
}

func main() {
	//播种随机数种子
	rand.Seed(time.Now().UnixNano())
	//quit := make(chan bool)		//用于关闭主go程的channel
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go readGo(ch, i+1)
	}
	for i := 0; i < 5; i++ {
		go writeGo(ch, i+1)
	}

	for{
		;
	}
}
