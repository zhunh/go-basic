package go_basic

import (
	"crypto/rand"
	"fmt"
	"math/big"
)


func main() {
	var numArr[10] *big.Int
	for i:=0;i<10;i++{
		numArr[i], _ = rand.Int(rand.Reader, big.NewInt(10))
	}

	fmt.Println(numArr)
}
