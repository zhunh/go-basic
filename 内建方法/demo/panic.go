package demo

import (
	"fmt"
)

//抛出异常

func ReceivePanic()  {
	defer CoverPanic()
	//panic("I am panic.")
	panic(2)
	//panic(errors.New("I am an error"))
}

func CoverPanic()  {
	message := recover()
	switch message.(type) {
	case string:
		fmt.Println("string message:",message)
	case error:
		fmt.Println("error message:",message)
	default:
		fmt.Println("unknow panic:",message)
	}
}