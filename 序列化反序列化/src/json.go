package src

import (
	"encoding/json"
	"fmt"
)

//序列化struct 和 map

type Student struct{
	Name string `json:"name"`
	Age int`json:"age"`
}

func Serialize()  {
	cat := Student{Name:"hy",Age:19}
	re,err := json.Marshal(cat)
	if err != nil{
		fmt.Println(err)
		return
	}else {
		fmt.Println("Marshal json: ",string(re))
	}

}

func DeSerialize()  {
	str := `{"name":"jack","age":23}`
	stu := new(Student)
	err := json.Unmarshal([]byte(str),&stu)
	if err != nil{
		fmt.Println(err)
		return
	}else {
		fmt.Println("Unmarshal struct: ",stu)
	}

}
