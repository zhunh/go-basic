package main

import "fmt"
/*
多重继承
*/
type people struct {
	name string
	age int
	sex string
}

type hero struct {
	people  //匿名字段继承
	capable string
	id int
}

type avenger struct {
	alias string
}

type teenager struct {
	avenger
	hero	//匿名字段多继承
	major string
}

func main() {
	var spiderman teenager
	spiderman.name = "Peter"
	spiderman.age = 17
	spiderman.sex = "male"
	spiderman.capable = "hanging by web."
	spiderman.id = 1001
	spiderman.major = "计算机"
	spiderman.alias = "蜘蛛侠"
	fmt.Println(spiderman)
}