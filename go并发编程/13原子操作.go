package main

import (
	"fmt"
	"sync/atomic"
)
/*
原子操作，同步操作6种基础类型
功能上相当于加锁访问
具体实现依靠了底层硬件的支持，所以操作效率高于锁
*/
func main() {
	var a int64 = 123
	//原子操作一定成功
	atomic.StoreInt64(&a, 233)

	fmt.Println(a)
	//将a 的值换为555
	old := atomic.SwapInt64(&a, 555)
	fmt.Println(old, a)
	//确保旧值为555 的前提下，a 的值换为666
	ok := atomic.CompareAndSwapInt64(&a, 555, 666)
	if ok{
		fmt.Println(a)
	}
}
