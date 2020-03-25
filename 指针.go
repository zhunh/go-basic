package main

import "fmt"

func main()  {
	var arr [5]int = [5]int{1,2,3,4,5}

	//fmt.Printf("%p\n",&arr)
	//fmt.Printf("%p\n",&arr[0])

	//p:=&arr
	//fmt.Printf("%T\n",p)
	//定义指针指向数组
	var p *[5]int
	//将指针变量和数组建立联系
	p=&arr
	//fmt.Printf("%p\n",p)
	//fmt.Println(p)

	//可以通过指针间接操作数组
	fmt.Println(*p)
	fmt.Println((*p)[0])
	//直接使用指针[下标]操作数组元素
	fmt.Println(p[0])


}
