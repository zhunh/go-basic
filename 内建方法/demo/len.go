package demo

import "fmt"

//len -> string array slice map chan
//cap -> slice array chan
//close -> chan

func GetLen()  {
	mslice := make([]string, 2, 5)
	mslice[0] = "id-1"
	mslice[1] = "id-2"
	fmt.Println("mslice len:",len(mslice))
	fmt.Println("mslice cap:",cap(mslice))

}