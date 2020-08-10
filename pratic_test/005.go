package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
请编写一个方法，将字符串中的空格全部替换为“%20”。
假定该字符串有足够的空间存放新增的字符，并且知道字符串
的真实长度(小于等于1000)，同时保证字符串由【大小写的英文字母组成】。
给定一个string为原始的串，返回替换后的string。
*/
func replaceString(s string) (string, bool) {
	//长度判断
	if len([]rune(s)) > 1000 {
		fmt.Println("长度大于1000")
		return s, false
	}
	for _, v := range s {
		//确保字符串都为字母
		if string(v) != " " && unicode.IsLetter(v) == false {
			return s, false
		}
	}
	s = strings.Replace(s, " ", "%20", -1)
	return s, true
}
func main() {
	//fmt.Println(replaceString("df jkd  b"))
	str := "trt"
	fmt.Println(str[1:2])
}
