package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex04 sync.RWMutex		//锁只有一把，两个属性 r w

var value int		//定义全局变量，模拟共享数据

func readGo04(idx int)  {
	for{
		rwMutex04.RLock()		//以读模式枷锁
		num := value
		fmt.Printf("--%dth 读go程， 读出：%d\n", idx,num)
		rwMutex04.RUnlock()	//以读模式解锁
		time.Sleep(time.Second)
	}
}

func writeGo04(idx int)  {
	for{
		//生成随机数
		num := rand.Intn(100)

		rwMutex04.Lock()		//以写模式加锁
		value = num
		fmt.Printf("%dth 写go程，写入：%d\n", idx,num)
		time.Sleep(time.Millisecond * 300)
		rwMutex04.Unlock()
	}
}

func main() {
	//播种随机数种子
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		go readGo04(i+1)
	}
	for i := 0; i < 5; i++ {
		go writeGo04(i+1)
	}

	for{
		;
	}
}
