package utils

import (
	"github.com/go-redis/redis/v7"
	"time"
)

var Pool *redis.Client

func InitPool()  {
	Pool = redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr: "0.0.0.0:6379",
		Password:"",
		DB: 0,
		PoolSize:10,	//连接池大小
		MaxRetries: 3,	//最大重试次数
		IdleCheckFrequency: 10*time.Second, //空闲连接超时时间
	})
}
