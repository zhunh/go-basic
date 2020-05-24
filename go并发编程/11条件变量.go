package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	//监听条件
	var condition = false

	go func() {
		time.Sleep(time.Second*3)

		cond.L.Lock()
		//变更条件
		condition = true
		fmt.Println("[1]条件变更。")
		//cond.Signal()将条件变更通知给关注该条件的一条协程
		//cond.Broadcast()将条件变更通知给关注该条件的所有协程
		//cond.Signal()
		cond.Broadcast()
		fmt.Println("[1]条件变更已通知，继续后续处理。")
		cond.L.Unlock()
	}()
	//协程2监听condition 变化
	go func() {
		cond.L.Lock()
		for !condition{
			fmt.Println("[2]条件尚未变更，等待变更通知。")
			cond.Wait()
		}
		cond.L.Unlock()
		fmt.Println("[2]收到条件变更通知，开始处理。")
	}()

	cond.L.Lock()
	for !condition{
		fmt.Println("[main]条件尚未变更，等待变更通知。")
		cond.Wait()
	}
	cond.L.Unlock()

	fmt.Println("[main]收到条件变更通知，开始处理。")
	//time.Sleep(100*time.Second)
	fmt.Println("main over.")

}
