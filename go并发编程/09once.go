package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	Name string
	Alive bool
}

func Kill(p *Person)  {
	fmt.Println("刺杀",p)
	<-time.After(time.Second*1)
	p.Alive = false
	fmt.Printf("刺杀 %s成功~\n",p.Name)
}



func main() {
	//声明Once对象
	var once sync.Once
	var wg1 sync.WaitGroup

	bill := &Person{"比尔",true}
	//开辟10条协程刺杀比尔
	for i:=0; i<10; i++{
		wg1.Add(1)
		go func() {
			//保证只将比尔杀死一次
			once.Do(func() {
				Kill(bill)
			})
			wg1.Done()
		}()
	}

	wg1.Wait()
}
