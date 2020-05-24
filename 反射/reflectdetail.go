package main

import (
	"fmt"
	"reflect"
)

//通过反射修改 num int的值，修改 b box 的值

func test(b interface{})  {
	val := reflect.ValueOf(b)
	fmt.Printf("val kind=%v\n",val.Kind())
	val.Elem().SetInt(20)
}

func main() {
	var num int = 10
	test(&num)//传入指针
	fmt.Println("num=", num)
}
