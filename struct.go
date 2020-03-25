package go_basic

import "fmt"

//在函数外定义结构体，作用域是全局
//type 结构体名 struct{
//	结构体成员列表
//}

type Student struct {
	id   int
	name string
	sex  string
	age  int
	addr string
}

func main(){
	//通过结构体名 定义结构体变量
	var s Student
	//结构体变量名.成员名
	s.id = 101
	s.name = "zhang"
	s.sex = "男"
	s.age = 29
	s.addr = "江西赣州"

	fmt.Println(s)

	var s1 Student = Student{102,"杨幂","女",18,"深圳"}
	fmt.Println(s1)
	fmt.Println(s1.name)


	s2 := Student{age:20,name:"张艺兴",sex:"男",addr:"广东",id:205}
	fmt.Println(s2)
	//结构体变量赋值
	s3:=s2
	//在结构体中使用 == 可以对结构体成员进行比较操作
	if s3 == s2{
		fmt.Println("相同")
	}else {
		fmt.Println("不相同")
	}

	//结构体数组
	//var arr [3]Student
	//for i:=0;i<len(arr);i++{
	//	fmt.Scan(&arr[i].id,&arr[i].age,&arr[i].name,&arr[i].sex,&arr[i].addr)
	//	fmt.Println(arr[i])
	//}

	//结构体切片
	a := []Student{{107,"ccvb","male",19,"mars"},
		{108,"qqq","male",18,"mars"}}
	a = append(a,Student{109,"ccvb","male",19,"mars"})
	fmt.Println(a)
}
