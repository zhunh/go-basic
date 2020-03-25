package main

import "fmt"

func main() {
	for{
		var tmp string
		fmt.Scanln(&tmp)
		fmt.Println("你输入的是：",tmp)
	}
}
