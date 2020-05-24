package main

import (
	"fmt"
	"sync"
	"time"
)
/*使用10个协程同时频繁修改一个数据，演示何为“并发不安全”*/
func main3() {
	var wg sync.WaitGroup
	var money = 2000

	for i:=0; i<10; i++{
		wg.Add(1)
		go func() {
			for j:=0; j<1000; j++{
				money +=1
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(money)//结果理论为12000，实际却是小于12000的，这里体现了并发不安全
}

/*使用同步锁后可以安全访问*/
func main() {
	var mt sync.Mutex

	var wg sync.WaitGroup
	var money = 2000

	for i:=0; i<10; i++{
		wg.Add(1)
		go func(i int) {
			fmt.Printf("协程%d开始抢锁\n", i)
			/*
			抢锁
			一次只能被一个协程锁住
			其余想要想到这把锁的协程阻塞等待至前面的协程将锁释放
			1.抢到锁，向下执行
			2.没抢到，阻塞等待至前面的协程将锁释放
			*/
			mt.Lock()
			fmt.Printf("协程%d抢到锁\n", i)
			fmt.Printf("协程%d开始操作数据\n", i)
			for j:=0; j<1000; j++{
				money +=1
			}
			<-time.After(3*time.Second)//耗时操作，等待一秒
			/*
			解锁
			释放之后，其他抢这把锁的协程就会得到抢锁机会
			*/
			mt.Unlock()
			fmt.Printf("协程%d释放锁\n", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(money)//结果理论为12000，实际却是小于12000的，这里体现了并发不安全
}
