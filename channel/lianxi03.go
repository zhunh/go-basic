package main

import "fmt"

/*
判断8000以内的素数
*/

func pData(inChan chan int)  {
	for i:=1; i<80; i++{
		inChan<- i
	}
	close(inChan)
}

func gData(inChan chan int, primeChan chan int, exitChan chan bool)  {
	var flag bool//标识是不是素数
	for{
		num, ok := <-inChan
		if !ok{ //intChan 取不到数据了
			break
		}
		flag = true
	   	//判断num是不是素数
	   	for i:=2; i<num; i++{
	   		if num%i == 0{//说明该num不是素数
				flag = false
				break
			}
		}
	   	if flag{
	   		//将这个数就放入到primeChan
	   		primeChan<- num
		}
	}
	fmt.Println("有一个primeNum 协程因为取不到数据退出了。")
	//这里还不能关闭 primeChan ,向exitChan 写入true
	exitChan<- true
}

func main() {
	intChan := make(chan int, 1000)
	resChan := make(chan int, 2000)//存放素数
	exitChan := make(chan bool, 4) //标识退出的通道

	go pData(intChan)
	//开启四个协程， 从 intChan 中取出数据， 并判断是否为素数
	for i:=0;i<4; i++{
		go gData(intChan, resChan, exitChan)
	}

	go func() {
		for i:=0;i<4;i++{//这里会读四次吗？
			<-exitChan
		}
		//当我们从exitChan取出4个结果
		close(resChan)
	}()

	for{
		res, ok := <-resChan
		if !ok{
			break
		}
		//将结果输出
		fmt.Printf("素数=%d\n",res)
	}
	fmt.Println("主线程退出。")
}
