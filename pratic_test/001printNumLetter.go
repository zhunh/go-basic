package main

import (
	"fmt"
	"strings"
	"sync"
)

//使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
//12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
func printNum(i int) {
	fmt.Println(i)
}

func main() {
	var wg sync.WaitGroup
	chanNum, chanLetter := make(chan bool), make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-chanNum:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				chanLetter <- true
				break
			default:
				break
			}
		}
	}()

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		s := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-chanLetter:
				if i >= strings.Count(s, "")-1 {
					wg.Done()
					return
				}
				fmt.Print(s[i : i+1])
				i++
				fmt.Print(s[i : i+1])
				i++
				chanNum <- true
				break
			default:
				break
			}
		}
		wg.Done()
	}(&wg)
	chanNum <- true
	wg.Wait()
}
