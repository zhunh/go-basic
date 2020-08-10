package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)

//定义Redis连接池
var client *redis.Client

//初始化redis连接池
func init() {
	client = redis.NewClient(&redis.Options{
		Addr: "0.0.0.0:6379",
		Password:"",
		DB: 0,
		PoolSize:10,
		MaxRetries: 3,
		IdleCheckFrequency: 10*time.Second,
	})

	pong, err := client.Ping().Result()
	if err == redis.Nil {
		fmt.Println("Redis异常")
	} else if err != nil {
		fmt.Println("失败:", err)
	} else {
		fmt.Println("连接正常：", pong)
	}

	fmt.Println(client.PoolStats())
}

func main()  {
	//增加
	putRes, err := client.HSet("zhuz", "ss", "ji").Result()
	if err != nil{
		fmt.Println("err=",err)
	}
	fmt.Printf("putRes=%v\n", putRes)

	//查询
	getRes, err := client.HGet("users","100").Result()
	if err != nil{
		fmt.Println("err=",err)
	}
	fmt.Println("getRes=", getRes)
}