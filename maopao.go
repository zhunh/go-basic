package go_basic

import "fmt"

func bubbleSort(arr []int) []int{
	for i:=0; i<len(arr)-1; i++{
		for j:=0; j<len(arr)-i-1; j++{
			if arr[j]>arr[j+1]{
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
	return arr
}

func main(){
	var test = []int{3,2,6,4,1}
	fmt.Println(test)
	bubbleSort(test)
}
