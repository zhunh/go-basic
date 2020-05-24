package main

import "log"

type ConnLimiter struct {
	concurrentConn int
	bucket chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	//桶满了， 没拿到token， 返回false
	if len(cl.bucket) >= cl.concurrentConn{
		log.Printf("Reached the rate limitation.")
		return false
	}
	//桶没满， 拿到token， 返回true
	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn()  {
	c := <- cl.bucket
	log.Printf("New connect coming: %d", c)
}