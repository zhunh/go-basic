package main

import "fmt"

//在函数内部修改值，不会影响函数外的值
func swap1(a, b int) {
	temp := b
	b = a
	a = temp
}

//通过返回值实现数值交换
func swap2(a, b int) (c, d int) {
	c, d = b, a
	return
}

//指针操作改变函数外的值
func swap3(a, b *int) {
	*a, *b = *b, *a
}
func main() {
	i, j := 1, 2

	swap3(&i, &j)
	fmt.Println(i, j)

}
