package main

import "fmt"

type student struct {
	name string
	sex string
	age int
	yw int
	sx int
	yy int
}

func (s *student)sayHi(){
	fmt.Printf("大家好，我叫%s,\n今年%d岁了,是%s同学。\n",s.name,s.age,s.sex)
}

func (s *student)printScore(){
	fmt.Println("总分：",s.yw+s.sx+s.yy)
	fmt.Println("平均分数：",(s.yw+s.sx+s.yy)/3)
}

func main()  {
	var xm = student{"小明","男",21,78,89,99}
	xm.sayHi()
	xm.printScore()
}