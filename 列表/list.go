package main

import (
	"container/list"
	"fmt"
)

//双链表
func main() {
	var li list.List
	//插入尾部
	li.PushBack("first")
	//插入头部
	li.PushFront(00)

	el := li.PushBack("ko")
	//在el 元素之后插入 ng
	li.InsertAfter("ng", el)
	//在 el 元素之前插入 bv
	li.InsertBefore("bv", el)

	fmt.Println(el)
	//删除
	li.Remove(el)

	//遍历
	for i := li.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
