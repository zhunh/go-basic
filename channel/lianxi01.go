package main

import (
	"fmt"
)

func writeData(intChan chan int)  {
	for i:=0; i<50; i++{
		intChan<- i
		fmt.Println("写入数据：", i)
	}
	//close(intChan)
}

func readData(intChan chan int, qChan chan bool)  {
	//for i := range intChan{
	//	fmt.Println("读到数据：", i)
	//}
	//qChan<- true
	for{
		v, ok := <-intChan
		if !ok{
			break
		}
		fmt.Printf("读到数据 %d\n", v)
	}
	//读取完成
	qChan<- true
	close(qChan)
}

func main() {
	var c = make(chan int, 50)
	var exitChan = make(chan bool, 1)
	go writeData(c)
	go readData(c, exitChan)
	//select {
	//case c := <-q:
	//	if c == true{
	//		fmt.Println("退出。")
	//	}
	//default:
	//}

	//for {
	//	var f = <-exitChan
	//	if f{
	//		break
	//	}
	//}
	for{
		_, ok := <-exitChan
		if ok{
			break
		}
	}
}
