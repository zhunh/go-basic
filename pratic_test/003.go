package main

import "fmt"

func reversString(s string) string {
	str := []rune(s)
	len := len(str)
	mid := len / 2
	for i := 0; i < mid; i++ {
		str[i], str[len-i-1] = str[len-i-1], str[i]
	}
	return string(str)
}

func main() {
	fmt.Println(reversString("hsdkfdf"))
}
