package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
有一个机器人，给一串指令，L左转 R右转，F前进一步，B后退一步，
问最后机器人的坐标，最开始，机器人位于 0 0，方向为正Y。 可以
输入重复指令n ： 比如 R2(LF) 这个等于指令 RLFLF。
问最后机器人的坐标是多少？
复杂指令 R2(L2(F)3(gf)) 、 R2(B2(LF)BF2(BF))FBF
*/
//将指令去除括号
func handleCmd(s string) string {
	//1.找到最左边的右括号
	rIndex := strings.Index(s, ")")
	//2.如果没有找到右括号，返回s
	if rIndex == -1 {
		fmt.Println(s)
		return s
	}
	//3. 字符串s以找到的最左边右括号为分界点，找s左边最后出现的左括号
	//[即与s最左边出现的右括号匹配的那个左括号]
	lIndex := strings.LastIndex(s[:rIndex], "(")
	//4.找到该对括号前面的重复次数  eg:得到 2(BF) 中的 2
	repeatCount, _ := strconv.Atoi(string(s[lIndex-1]))
	//5.重复括号里面的字符切片 eg:得到 2(BF) 中的 BF
	sliceStr := s[lIndex+1 : rIndex]
	//6.组装 eg: BF => BFBF
	replaceStr := strings.Repeat(sliceStr, repeatCount)
	//7.替换  将2(BF) 替换为 BFBF
	s = strings.Replace(s, s[lIndex-1:rIndex+1], replaceStr, -1)
	//8.递归处理下一对括号
	return handleCmd(s)
}

func main() {
		var x, y int = 0, 0
		var cmd string
		for {
			fmt.Scanln(&cmd)
			c := handleCmd(cmd)
			for _,v := range c{
				switch string(v) {
				case "L":
					x--
				case "R":
					x++
				case "F":
					y++
				case "B":
					y--
				default:
					fmt.Println("wrong cmd.")
				}
			}
			fmt.Printf("机器人现在的位置为(%d,%d)\n", x, y)
		}
}
