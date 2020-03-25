package go_basic

import "fmt"
//必须写元素个数的为数组，不写的为切片，切片可以扩展
func main() {
	//数组定义 var 数组名 [元素个数]数组类型
	//var s [5]int
	//切片定义 var 切片名 []数据类型
	//var s []int
	s := make([]int,5)
	//通过下标为切片赋值
	s[0] = 123
	s[1] = 456
	s[2] = 789
	s[3] = 892
	s[4] = 70
	//通过append为切片添加元素
	s=append(s,12,34,56)
	//容量每次扩展为上次的倍数
	fmt.Println("容量：",cap(s))
	fmt.Println("长度：",len(s))

	//遍历
	//for i:=0; i<len(s);i++{
	//	fmt.Println(s[i])
	//}
	for i,v:=range s{
		fmt.Println(i,v)
	}
	//切片截取
	//切片名[起始位置:结束位置] 不包含结束位置
	slice := s[0:2]
	fmt.Println(cap(slice))
	fmt.Println(slice)
	//切片名[起始位置:结束位置:指定容量] 容量大于等于len小于等于cap
	slice2 := s[0:2:5]
	fmt.Println(cap(slice2))
	fmt.Println(slice2)
}
