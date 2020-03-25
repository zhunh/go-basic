package main

import "fmt"

type person2 struct {
	name string
	sex string
	age int
}

type student2 struct {
	person2
	score int
}

type teacher2 struct {
	person2
	subject string
}
//定义接口
type Xer interface {
	//方法 函数声明 X没有具体实现，具体实现根据对象的不同，实现方式也不同
	PrintInfo()
}

func(t teacher2) PrintInfo(){
	fmt.Printf("大家好，我是老师%s，今年%d岁，是%s生,教的学科是%s\n",t.name,t.age,t.sex,t.subject)
}

func(s student2) PrintInfo(){
	fmt.Printf("大家好，我是学生%s，今年%d岁，是%s生，得分%d。\n",s.name,s.age,s.sex,s.score)
}

func main() {
	var stu = student2{person2{"jack","男",19},98}
	var tea = teacher2{person2{"rose","女",23},"英语"}
	var tt Xer
	tt = &stu
	tt.PrintInfo()

	tt = &tea
	tt.PrintInfo()
}