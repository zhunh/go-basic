package go_basic

import "fmt"

//map定义
//var 变量名 map[keyType]valueType

func main(){
	//创建一个空的map
	//var m map[int]string
	//记住这种定义方式，初始定义容量为1，但是map的长度会自动扩容
	//map中的数据是无序存储的
	m := make(map[int]string,1)
	m[199]="haha"
	m[198]="haha" //map中值允许重复
	m[1]="zhunianhong"
	m[23]="Curry"
	fmt.Println(m)
	fmt.Println(m[1])

	//遍历
	for k,v := range  m{
		fmt.Println(k,v)
	}

	//输出判断
	v,ok := m[5]
	if ok{
		fmt.Println(v)
	}else {
		fmt.Println("key不存在。")
	}

	//删除map中的元素
	m2 := map[int]string{101:"图书馆",46:"主教",90:"实验数据"}
	fmt.Println(m2)

	delete(m2,46)
	delete(m2,46)  //delete删除的时候，如果key不存在 不会报错

	fmt.Println(m2)

	//map作为函数参数是地址传递（引用传递）
	demo(m2)
	fmt.Println(m2)
}

func demo(m map[int]string)  {
	m[789]="食堂"
	m[20000]="操场"
	fmt.Println(len(m))
}
