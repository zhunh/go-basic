package main

import (
	"fmt"
	"os"
)

func main() {
	hostName, _ := os.Hostname()
	fmt.Println("hostName:", hostName)
	fmt.Println("系统内存页的尺寸",os.Getpagesize())
	//fmt.Println("环境变量：",os.Environ())
	fmt.Println(os.Geteuid())
}

