package gorotine

import (
	"fmt"
	"time"
)

func Loop1()  {
	for i:=0;i<10;i++{
		fmt.Println("L1-",i)
		time.Sleep(time.Microsecond * 10)
	}
}

func Loop2()  {
	for i:=0;i<10;i++{
		fmt.Println("L2-",i)
		time.Sleep(time.Second * 1)
	}
}
