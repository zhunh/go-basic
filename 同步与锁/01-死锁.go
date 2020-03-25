package main

/*
	1.单go程自己死锁
		channel必须在两个以上的go程中通信，否则死锁！
	2.go程间channel访问顺序导致死锁
		使用channel一端读（写），要保证另一端写（读）操作，同时有机会执行。否则死锁。
	3.多go程，多channel交叉死锁
*/
//死锁1
/*func main() {
	ch := make(chan int)
	ch <- 123
	num := <-ch
	fmt.Println("num = ", num)
}*/

//死锁2
/*func main() {
	ch:=make(chan int)

	num := <- ch
	fmt.Println("num = ", num)

	go func() {
		ch <- 123
	}()
}*/

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go func() {
		for {
			select {
			case num := <- ch1:
				ch2 <- num
			}
		}
	}()

	for {
		select {
			case num := <-ch2:
				ch1 <- num
		}
	}
}
