package main

import (
	"fmt"
	"going/go-basic/并发/gorotine"
	"time"
)

func main()  {
	//顺序执行
	//gorotine.Loop1()
	//gorotine.Loop2()

	//并行执行
	//fmt.Printf("cupNum:%d",runtime.NumCPU())
	//控制CPU使用数量
	//runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	//go gorotine.Loop1()
	//go gorotine.Loop2()
	//time.Sleep(time.Second * 30)

	//测试协程通信
	//go gorotine.Send()
	//go gorotine.Recv()
	//time.Sleep(time.Second * 30)

	//测试协程同步
	gorotine.Read()
	go gorotine.Write()
	gorotine.WG.Wait()
	fmt.Println("All done")  //主线程将一直等到Write线程结束才打印这句
	time.Sleep(time.Second * 30)
}
