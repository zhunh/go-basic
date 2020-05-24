package main

import (
	"fmt"
	"reflect"
)

func dd(i interface{})  {
	t := reflect.TypeOf(i)
	fmt.Printf("Type=%v\n",t)
	v := reflect.ValueOf(i)
	fmt.Printf("Value=%v\n",v)

	fmt.Printf("kind=%v\n", t.Kind())

	iv := v.Interface()
	fmt.Printf("iv=%v, iv type=%T\n", iv, iv)
	n := iv.(float64)
	fmt.Printf("n=%f\n",n)
}

func main() {
	var v float64 = 1.2
	fmt.Println("v=",v)
	dd(v)
}
