package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	//"os"
)

type Whisky struct {
	Whisky []string
}

func main() {
	//pwd,_ := os.Getwd()
	fileInfoList,err := ioutil.ReadDir(`/Users/znh/Desktop/wine_demo/金酒`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(fileInfoList))
	var nameList []string
	for i:= range fileInfoList{
		//nameList[i] = fileInfoList[i].Name()
		nameList = append(nameList, fileInfoList[i].Name()[0:len(fileInfoList[i].Name())-4])
	}
	fmt.Println(len(nameList))

	tmp := Whisky{nameList}
	//序列化
	b, err := json.Marshal(tmp)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(b))
	//写入json文件
	err = ioutil.WriteFile("nameList.json", b, os.ModeAppend)
	if err != nil{
		return
	}
}
