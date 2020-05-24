package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}){
	//通过反射获取传入变量的 type， kind，值
	//1.获取reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("Type=",rTyp)
	//2.获取reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Println("Value=", rValue)

	//3.将 rValue转成 interface{}
	iv := rValue.Interface()
	//将 interface{} 通过类型断言转成需要的类型
	num2 := iv.(int)
	fmt.Println("num2=", num2)
}
type student struct {
	name string
	age int
}
//演示对结构体的反射
func reflectTest02(b interface{})  {
	//通过反射获取传入变量的 type， kind，值
	//1.获取reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("Type=",rTyp)
	//2.获取reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Println("Value=", rValue)

	//3.获取变量对应的Kind
	vKind := rValue.Kind()
	fmt.Printf("vkind=%v\n",vKind)

	iv := rValue.Interface()
	fmt.Printf("iv=%v, iv type=%T\n",iv,iv)

	stu, ok := iv.(student)
	if ok{
		fmt.Printf("stu.name=%v\n",stu.name)
	}
}

func main() {
	//编写反射演示案例
	//var n int = 100
	//reflectTest(n)
	//结构体反射
	stu := student{"zhu", 19}
	reflectTest02(stu)
}
