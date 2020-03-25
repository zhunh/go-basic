package demo

import (
	"fmt"
	"reflect"
)

/*
make返回引用类型
*/

//创建切片
func MakeSlice()  {
	mSlice := make([]string,3)
	mSlice[0] = "dog"
	mSlice[1] = "cat"
	mSlice[2] = "lion"
	fmt.Println(mSlice)
}
//创建map
func MakeMap()  {
	mMap := make(map[int]string)
	mMap[10] = "math"
	mMap[12] ="physics"
	fmt.Println(mMap)
	fmt.Println(reflect.TypeOf(mMap))
}
//创建通道 没有缓存的chan
func MakeChan()  {
	mChan := make(chan int)
	close(mChan)
}