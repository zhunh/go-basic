package demo

import "fmt"

//删除map中的数据

func DeleteFromMap()  {
	mMap := make(map[int]string)
	mMap[101] = "id-201"
	mMap[102] = "id-202"
	delete(mMap,102)
	fmt.Println(mMap)
}
