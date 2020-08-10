package main

import "fmt"

func main() {
	if true {
		defer fmt.Println("1")
	} else {
		defer fmt.Println("2")
	}
	fmt.Println("3")
}

//接过为 3 1