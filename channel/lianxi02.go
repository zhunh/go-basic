package main

import (
	"fmt"
)

func putData(intChan chan int)  {
	for i:=1; i<2001;i++{
		intChan<- i
		fmt.Println("写入数据：",i)
	}
	close(intChan)
}

func pullData(intChan chan int, resChan chan int, exitChan chan bool) {
	for{
		v, ok := <-intChan
		if !ok{
			break
		}
		resChan<- sum(v)
		fmt.Printf("读出%d，和为%d\n",v,sum(v))
	}
	exitChan<- true
	close(exitChan)
}

func sum(n int) int {
	var h int = 0
	for i:=0;i<n;i++{
		h+=i
	}
	return h
}

func main() {
	var intChan = make(chan int, 2000)
	var resChan = make(chan int, 2000)
	var exitChan = make(chan bool, 1)
	go putData(intChan)

	for i:=0;i<8;i++{
		go pullData(intChan,resChan,exitChan)
	}

	for{
		_, ok := <-exitChan
		if ok{
			break
		}
	}
}
