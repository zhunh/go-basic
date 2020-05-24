package main

import (
	"fmt"
	"time"
)

/*
两个协程，一个发生panic，会导致整个程序崩溃
我们在协程中使用 recover 来捕获异常，使程序不至于所有的协程崩溃
*/

func sayHello()  {
	for i:=0; i<10; i++{
		time.Sleep(time.Second)
		fmt.Println("hello, word")
	}
}

func test()  {
	//这里使用 defer + recover
	defer func() {
		//捕获test抛出的panic
		if err := recover();err != nil{
			fmt.Println("test() 发生错误",err)
		}
	}()
	var myMap map[int]string
	myMap[0]="golang" //这里直接调用myMap 会有错误
}

func main() {
	go sayHello()
	go test()

	for i:=0; i<10; i++{
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}
}
