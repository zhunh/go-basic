package main

import (
	"fmt"
	"reflect"
)

//使用反射来遍历结构体的属性和方法

type hero struct {
	Name string `json:"heroname"`
	tech string `json:"superpower"`
	age int `json:"heroage"`
}

func (h hero) Say()  {
	fmt.Printf("hi i'm %v, i can %v, and i'm %v years old.\n", h.Name, h.tech, h.age)
}

func (h hero) Add(a, b int)  {
	fmt.Printf("%d + %d = %d\n", a, b, a + b)
}

func bb(i interface{})  {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	kd := v.Kind()
	//判断传入类型是否为struct
	if kd != reflect.Ptr && v.Elem().Kind() == reflect.Struct{
		fmt.Println("expect struct param")
		return
	}
	//获取该结构体有多少个字段
	fieldNum := v.NumField()
	fmt.Printf("total %d field in this struct.\n", fieldNum)
	//遍历该结构体的各个字段
	for x:=0; x<fieldNum; x++{
		fmt.Printf("field %d = %v, tag= %v\n", x, v.Field(x), t.Field(x).Tag.Get("json"))
	}

	//获取该结构体有多少个方法
	methodNum := v.NumMethod()

	fmt.Printf("total %d methods.\n",methodNum)
	//调用结构体的第一个方法，多个方法排序按照方法名 ASCII 码排序
	v.Method(1).Call(nil)
	//调用结构体的第一个方法 add
	var params  []reflect.Value
	params = append(params, reflect.ValueOf(3))
	params = append(params, reflect.ValueOf(6))
	v.Method(0).Call(params)

}

func vv(n interface{})  {
	v := reflect.ValueOf(n)
	//修改字段值
	fNum := v.Elem().NumField()
	fmt.Println(fNum)
	v.Elem().Field(0).SetString("iroman")
	for i:=0; i<fNum; i++{
		fmt.Printf("%d %v\n",i, v.Elem().Field(i).Kind())
	}
}

func main() {
	var spideman hero = hero{
		"spideman",
		"hanging",
		17,
	}
	spideman.Say()
	//bb(spideman)
	vv(&spideman)
	fmt.Print(spideman)
}
