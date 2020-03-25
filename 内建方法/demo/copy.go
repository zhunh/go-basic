package demo

import "fmt"

//拷贝切片
func CopyToSlice()  {
	mSliceDst := make([]string,3)
	mSliceDst[0] = "id-dst-1"
	mSliceDst[1] = "id-dst-2"
	mSliceDst[2] = "id-dst-3"

	mSliceSrc := make([]string,2)
	mSliceSrc[0] = "id-src-1"
	mSliceSrc[1] = "id-src-2"

	copy(mSliceDst,mSliceSrc)
	fmt.Println(mSliceDst)
}