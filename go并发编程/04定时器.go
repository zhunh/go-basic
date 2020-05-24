package main

import (
	"fmt"
	"os"
	"time"
)

func main1() {
	timer := time.NewTimer(3*time.Second)
	fmt.Println("计时开始", time.Now())
	x := <-timer.C
	fmt.Println("引爆于：", x)
	fmt.Println("我宣布。。今天去春游！！")
}
//定时器终止
func main2() {
	timer := time.NewTimer(3*time.Second)
	fmt.Println("炸弹计时开始", time.Now())
	go func() {
		time.Sleep(time.Second*2)//两秒之后拆除，赶在三秒之前
		ok := timer.Stop()
		if ok{
			fmt.Println("炸弹已拆除")
			os.Exit(0)
		}
	}()
	x := <-timer.C
	fmt.Println("引爆于：", x)
}
//定时器重置
func main() {
	timer := time.NewTimer(5*time.Second)
	fmt.Println("炸弹计时开始", time.Now())

	time.Sleep(time.Second)
	timer.Reset(time.Second*2)//以当前时间为基准重置为2秒

	x := <-timer.C
	fmt.Println("炸弹引爆于：", x)
}