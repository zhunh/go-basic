package demo

import (
	"fmt"
	"reflect"
)

/*
	内存指令
	返回传入类型的指针地址
*/

func NewMap()  {
	mMap := new(map[int]string)
	fmt.Println(reflect.TypeOf(mMap))
}