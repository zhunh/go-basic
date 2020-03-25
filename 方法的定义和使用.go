package main

import "fmt"

/*
type的两个作用：
	1.定义函数类型
	2.为已存在的数据类型起别名
*/
type Int int  //给int起个别名Int
//为Int类型定义add方法
//格式：func (方法接受者)方法名（参数列表）返回值类型
func (a Int)add(b Int) Int {
	return a+b
}

type Stu struct {
	name string
	age int
	sex string
}
//值传递
func (s Stu)printInfo(){
	fmt.Println(s.name)
	fmt.Println(s.sex)
	fmt.Println(s.age)
}
//引用传递
func (s *Stu) editInfo(name string,sex string,age int) {
	//结构体指针可以直接调用结构体成员
	s.name=name
	s.sex=sex
	s.age=age
}

func main(){
	//根据数据类型 绑定方法
	var a Int =10
	value := a.add(20)
	fmt.Println(value)

	var s Stu = Stu{name:"章帅",sex:"男",age:19}
	//对象.方法 调用
	s.printInfo()
	s.editInfo("jack","male",22)
	s.printInfo() //注意这里打印的还是原来的数据  说明参数是值传递 应改为引用传递==改为 *Stu
}