package main

import (
	"fmt"
	"sync"
	"time"
)
/*
读写锁
多路只读
一路只写
*/
func main101() {
	var rwm sync.RWMutex
	//锁定为写模式，一路只写
	rwm.Lock()
	//释放读写锁
	rwm.Unlock()

	//锁定为读模式，多路只读
	rwm.RLock()
	//释放读写锁
	rwm.RUnlock()
}

/*
数据库的一读多写
Read方法设定为多路读写
Write方法设定为单路只写
创建5读5写10条协程，观察读写锁效果
*/
func main()  {
	var wg sync.WaitGroup
	//声明读写锁
	var rwm sync.RWMutex

	for i:=0;i<5;i++{
		//读协程
		wg.Add(1)
		go func() {
			//锁定为只读模式，允许多个协程同时抢到读锁，此时不允许写
			rwm.RLock()
			fmt.Println("Read协程正在读取数据.")
			<-time.After(3*time.Second)//模拟耗时操作
			rwm.RUnlock()

			wg.Done()
		}()

		wg.Add(1)
		go func() {
			//锁定为写模式，只允许一条协程抢到锁
			rwm.Lock()
			fmt.Println("Write协程正在写数据.")
			<-time.After(3*time.Second)//模拟耗时操作
			rwm.Unlock()

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("main over.")
}