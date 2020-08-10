package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type peer struct {
	id             int //节点标识
	CreditVal      int //信用值
	consensusTimes int //参与共识次数
}

//生成给定数量peer
func NewPeerArr(num int) []*peer {
	var arr []*peer
	rand.Seed(time.Now().Unix())
	for i := 1; i <= num; i++ {
		p := &peer{
			id:        i,
			CreditVal: rand.Intn(100),
		}
		arr = append(arr, p)
	}

	return arr
}

//打印一个节点信息
func printPeer(p *peer, info ...interface{}) {
	fmt.Printf("%v:[节点：%d，信用值：%d，共识次数：%d]\n", info, p.id, p.CreditVal, p.consensusTimes)
}

//打印一组节点信息
func printPeerArr(p []*peer) {
	for i := 0; i < len(p); i++ {
		printPeer(p[i], i)
	}
}

//扫描并提取共识次数为零的节点
func splitPeer(pa []*peer) ([]*peer, []*peer) {
	var hasConsensus, noConsensus []*peer
	for i := 0; i < len(pa); i++ {
		if pa[i].consensusTimes > 0 {
			hasConsensus = append(hasConsensus, pa[i])
		} else {
			noConsensus = append(noConsensus, pa[i])
		}
	}
	return hasConsensus, noConsensus
}

//构建 对一组节点按信用值排序 的方法
type ByCredit []*peer

func (bc ByCredit) Len() int {
	return len(bc)
}
func (bc ByCredit) Swap(i, j int) {
	bc[i], bc[j] = bc[j], bc[i]
}
func (bc ByCredit) Less(i, j int) bool {
	return bc[i].CreditVal < bc[j].CreditVal
}

//按节点信用值排序
func sortByCredit(p []*peer) {
	sort.Sort(ByCredit(p))
}

//选择记账节点
func chooseLeader(p []*peer) *peer {
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(p))
	p[index].consensusTimes += 1
	return p[index]
}
func consensusExec(p []*peer) {
	consensusIndex := 0
	for {
		consensusIndex += 1
		_, noCa := splitPeer(p)
		if len(noCa) != 0 {
			leader := chooseLeader(noCa)
			printPeer(leader, fmt.Sprintf("第%d轮共识选举得到", consensusIndex))
		} else {
			//先排序
			//sortByCredit(p)
			break
		}

		time.Sleep(time.Second * 3)
	}
}
func main() {
	//1.创建一定数量的节点,并初始化信用值为零
	p := NewPeerArr(100)
	printPeerArr(p)
	//2.模拟共识过程
	consensusExec(p)

	//sortByCredit(p)
	//printPeerArr(p)
}
