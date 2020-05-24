package main

import (
	"fmt"
	"time"
)

func main()  {
	n := time.Now()

	t1 := n.Format("2006-01-02 15:04:05")
	t2 := n.String()

	fmt.Println(n.Unix())//时间戳
	fmt.Println(n.UnixNano())//时间戳，精确
	fmt.Println(t1)
	fmt.Println(t2)

	currentTime := "2020-06-01 12:04:01"
	//解析时间到time类型，UTC时区
	u, _ := time.Parse("2006-01-02 15:04:05",currentTime)
	fmt.Println(u)
}
