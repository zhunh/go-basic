package main

import (
	"fmt"
	"time"
)

func main() {
	var tickerStoped = false
	ticker := time.NewTicker(1*time.Second)

	go func() {
		time.Sleep(time.Second*9)
		ticker.Stop()
		tickerStoped = true
	}()

	for{
		if tickerStoped{
			break
		}
		//阻塞一秒
		<-ticker.C
		fmt.Println("我要去浪~")
	}
	fmt.Println("main over.")
}
