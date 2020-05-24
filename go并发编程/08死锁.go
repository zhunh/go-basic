package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
//案例，AB相互要求对方先发红包自己再发
func main102() {
	cha := make(chan int)
	chb := make(chan int)
	//a
	wg.Add(1)
	go func() {
		<-cha 		//a先从自己账户读数据，模拟b先发了红包
		chb<-123	//然后再给b发红包
		wg.Done()
	}()

	//b
	wg.Add(1)
	go func() {
		<-chb
		cha<-12
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("main over.")
}

//案例，读写互斥死锁
func main()  {
	var rwm sync.RWMutex
	ch := make(chan int)

	wg.Add(1)
	go func() {
		//锁定为只读，排斥写入协程
		rwm.RLock()

		//需要写入协程，没人写就永远读不出来
		x:=<-ch
		fmt.Println("读到:",x)
		rwm.RUnlock()

		wg.Done()
	}()

	wg.Add(1)
	go func() {
		//锁定为只写，排斥读取协程
		rwm.Lock()
		//需要读取协程
		ch<- 12
		fmt.Println("写入数据。")
		rwm.Unlock()

		wg.Done()
	}()

	wg.Wait()
}