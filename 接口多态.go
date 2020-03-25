package main

import "fmt"

type dog struct {
	name string
	eat string
	sing string
}

func (d dog)fan()  {
	fmt.Printf("我是小狗%s,我吃%s,我会唱%s。\n",d.name,d.eat,d.sing)
}

type cat struct {
	name string
	eat string
	dance string
}

func (c cat)fan()  {
	fmt.Printf("我是小猫%s,我吃%s,我会跳舞%s.\n",c.name,c.eat,c.dance)
}
//接口定义
type fer interface {
	fan()
}
//多态实现
//多态是将接口类型作为参数，多态实现接口的统一处理
func SelfInfo(f fer)  {
	f.fan()
}

func main() {
	var f fer
	f=&dog{"hh","banana","小苹果"}
	//f.SelfInfo()
	SelfInfo(f)
	f=&cat{"jj","fish","华尔兹"}
	//f.SelfInfo()
	SelfInfo(f)
}
