package demo

import "fmt"

//slice -> append & copy

func AppendElementToSlice()  {
	mIDSlice := make([]string,2)
	mIDSlice[0] = "id-1"
	mIDSlice[1] = "id-2"
	fmt.Println("len=",len(mIDSlice))
	fmt.Println("cap=",cap(mIDSlice))
	mIDSlice = append(mIDSlice, "id-3")
	fmt.Println(mIDSlice)
	fmt.Println("after len=",len(mIDSlice))
	fmt.Println("after cap=",cap(mIDSlice))

}