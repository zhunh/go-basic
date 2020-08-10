package main

import (
	"crypto/sha256"
	"fmt"
)

func crypto(data []byte) string {
	sha := sha256.New()
	sha.Write(data)
	cptStr := fmt.Sprintf("%x", sha.Sum(nil))
	return cptStr
}

func main() {
	cptStr := crypto([]byte("first data block。"))
	fmt.Println(cptStr)
}
