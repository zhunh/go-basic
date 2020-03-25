package main

import (
	"errors"
	"fmt"
)

var erro001 = errors.New("零不能作除数")

func div(a,b int) (int,error){
	//判断除数为零的情况
	if b == 0{
		return 0,erro001
	}
	return a/b, nil
}

func main(){
	fmt.Println(div(3,0))
}
