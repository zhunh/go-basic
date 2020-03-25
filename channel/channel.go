package main

import "fmt"

func testChannel(c chan int, input int)  {
	fmt.Println("进入testChannel")
	//写入数据
	c <- input
	fmt.Println("写入数据：",input)
}

func testChannel2(c chan int)  {
	fmt.Println("进入testChannel2")
	//从channel中读取数据
	data := <-c
	fmt.Println("读出数据：",data)
}

func main(){
	c := make(chan int,2)
	go testChannel(c,889)
	go testChannel(c,880)
	go testChannel(c,230)
	//go testChannel2(c)

	//main协程 向channel中写入数据
	for{
		;
	}
}
