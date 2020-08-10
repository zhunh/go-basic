package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//第⼀步：打开数据库,格式是 ⽤户名：密码@/数据库名称？编码⽅式
	conn, err := sql.Open("mysql","root:19950202@tcp(localhost:3306)/video_server?charset=utf8")
	if err != nil{
		panic(err.Error())
	}
	//关闭数据库
	defer conn.Close()
	fmt.Println("连接成功")
}
