package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	服务器过载控制
	监听最大客户端连接数
	服务端协程：只要服务器过载，就通知控制协程，并进入阻塞等待
	控制协程：消减客户端数量后通知服务端（过载预警已解除）
*/

const MAX_CONNECTIONS = 3

type Server struct {
	conns int 			//客户端连接数
	cond *sync.Cond
	chAlarm chan bool 	//是否预警
}

func NewServer() *Server {
	s := &Server{}
	s.cond = sync.NewCond(&sync.Mutex{})
	s.chAlarm = make(chan bool)
	return s
}
//服务端
func (s *Server)Server(){
	for{
		//加锁检查服务器是否过载
		s.cond.L.Lock()
		//监听是否过载
		if s.conns == MAX_CONNECTIONS{
			s.chAlarm<- true
			fmt.Println("服务器过载预警")
			//阻塞等待预警解除
			s.cond.Wait()
		}
		s.cond.L.Unlock()
		//接入客户端
		fmt.Println("开始接入客户端")
		s.conns++
		fmt.Println("客户端数量：",s.conns)
	}
}
//释放连接
func (s *Server) Release()  {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	<- s.chAlarm
	s.cond.L.Lock()
	s.conns -= (random.Intn(MAX_CONNECTIONS)+1) //使用随机数来模拟当前客户端连接数
	fmt.Println("[Release]客户端数量：", s.conns)
	s.cond.Signal()
	fmt.Println("[Release]过载预警解除")
	s.cond.L.Unlock()
}

func main() {
	s := NewServer()
	go s.Server()
	//模拟一次释放客户端连接
	//time.Sleep(1*time.Second)
	go s.Release()

	//time.Sleep(1*time.Second)
	s.Release()

	//time.Sleep(time.Second*2)
}
