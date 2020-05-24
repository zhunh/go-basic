package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
开辟三条协程
A每秒生成一个三位随机数
B输出该值的奇偶性
C输出奇数和偶数的总个数
当生成一个水仙花数时，程序结束
*/

func GenRandom(start, end int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return(start + r.Intn(end-start+1))
}

func isEven(num int) bool {
	return num%2 == 0
}

func isShui(n int) bool {
	a := n / 100
	b := n %100 / 10
	c := n % 10
	return (a*a*a + b*b*b + c*c*c) == n
}

func main() {
	//var wg sync.WaitGroup
	var chNum = make(chan int)
	var chValueEven = make(chan [2]interface{})
	var chQuit  = make(chan bool)

	//wg.Add(1)
	go func() {
		ticker := time.NewTicker(time.Millisecond * 1)
		for{
			random := GenRandom(100, 999)
			fmt.Println("生成了随机数:",random)
			chNum<- random
			<-ticker.C
		}
		//wg.Done()
	}()

	//wg.Add(1)
	go func() {
		for x := range chNum{
			even := isEven(x)
			chValueEven <- [2]interface{}{x, even}
			if even{
				fmt.Println(x, "是偶数")
			}else{
				fmt.Println(x, "是奇数")
			}

		}
		//wg.Done()
	}()

	//wg.Add(1)
	go func() {
		var oddCount, evenCount int
		for x := range chValueEven{
			if x[1].(bool){
				evenCount++
			}else {
				oddCount++
			}
			fmt.Printf("奇数个数：%d, 偶数个数：%d\n", oddCount, evenCount)

			//判断是和否水仙花数
			if isShui(x[0].(int)){
				fmt.Println("发现水仙花：",x)
				chQuit<-true
			}
		}
		//wg.Done()
	}()

	//wg.Wait()
	<-chQuit
	fmt.Println("main over.")
}
