package main

import "fmt"

type person1 struct {
	name string
	sex string
	age int
}

type student1 struct {
	person1
	score int
}

func(p person1) PrintInfo(){
	fmt.Printf("大家好，我是%s，今年%d岁，是%s生.\n",p.name,p.age,p.sex)
}

func(s student1) PrintInfo(){
	fmt.Printf("大家好，我是%s，今年%d岁，是%s生，得分%d。\n",s.name,s.age,s.sex,s.score)
}

func main() {
	var stu = student1{person1{"jack","男",19},98}
	//默认使用子类方法，采用就近原则
	//调用子类方法
	stu.PrintInfo()
	//调用父类方法
	stu.person1.PrintInfo()
}