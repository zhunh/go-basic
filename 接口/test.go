package main

import (
	"fmt"
)

type StoreService interface {
	GetDB() string
}

type Peer struct {
	HostName string
	SS       StoreService
}

func NewPeer(SS StoreService) *Peer {
	nb := &Peer{HostName: "test.com"}
	nb.SetNS(SS)
	return nb
}

func (p *Peer) SetNS(SS StoreService) {
	p.SS = SS
}

type Store struct {
	DB string
}

func NewStore() *Store {
	s := &Store{"chen"}
	return s
}

func (c *Store) GetDB() string {
	return c.DB
}

func main() {
	p := NewPeer(NewStore())

	fmt.Println(p.SS.GetDB())
}
