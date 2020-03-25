package main

import "fmt"

func sing(d chan int) {


	for i:=0;i<50;i++{
		sig := <- d
		fmt.Println("读出sig：",sig)
		//fmt.Println("----正在唱滑板鞋----")
	}

}

func dance(d chan int) {
	for i:=0;i<50;i++{
		d <- 89
		fmt.Println("写入数据")
	}
}

func main() {
	c := make(chan int)

	go dance(c)
	go sing(c)

	for{
		;
	}
}