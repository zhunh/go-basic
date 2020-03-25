package main

import "fmt"

//语法糖 ...  :=

func main()  {
	sugar1("A","B","X")
}
//可变长度参数 类型为string
func sugar1(values...string)  {
	for _, v := range values  {
		fmt.Println("---> v:",v)
	}
}
