package main

import "fmt"

//1.全局变量定义
var str string
var str2 = ""

//2.同一包内多个init函数
func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}
func init() {
	fmt.Println("init3")
}
func main() {
	fmt.Println(str)
	//3.字符串拼接只能用双引号
	//str3 := '123' + 'abc'
	str3 := "123" + "abc"
	fmt.Println(str3)

	s := fmt.Sprintf("abc%d", 123)
	fmt.Println(s)
	//4.类型转换
	type MyInt int
	var i MyInt = 1
	var j int = int(i)
	fmt.Println(j)

	var sum int = 17
	var count int = 5
	var mean float32
	mean = float32(sum)/float32(count)
	fmt.Println(mean)

}
func sd(a,b int) (int, error) {

}