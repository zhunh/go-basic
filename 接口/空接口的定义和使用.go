package main

import (
	"fmt"
)

func main() {
	//定义空接口类型，接收任意类型输入
	var i interface{}
	//接收整数
	i = 90
	fmt.Println(i)
	//接收字符串
	i = "haha"
	fmt.Println(i)
	//接收数组
	i = [3]int{2, 5, 8}
	fmt.Println(i)

	//定义空接口类型切片，使该切片可以存储不同类型的值
	var wan []interface{}
	wan = append(wan, 23, "book", [3]int{3, 4, 8})
	//打印该切片
	for _, v := range wan {
		//对v类型断言, 类型断言必须左边为接口类型
		if data, ok := v.(int); ok {
			fmt.Println("int类型：", data)
		} else if data, ok := v.(string); ok {
			fmt.Println("string类型：", data)
		} else if data, ok := v.([3]int); ok {
			fmt.Println("array类型：", data)
		}
		//fmt.Println(v)
	}
}
