package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
)

func main() {
	//返回一个在[0, max)区间服从均匀分布的随机值，如果max<=0则会panic
	n, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err == nil {
		fmt.Println("crypto随机int:", n, n.BitLen())
	}
	//返回一个具有指定字位数（二进制的位数）的数字，该数字具有很高可能性是质数。如果从rand读取时出错，或者bits<2会返回错误
	p, err := rand.Prime(rand.Reader, 5)
	if err == nil {
		fmt.Println("crypto随机二进制:", p)
	}
	//本函数是一个使用io.ReadFull调用Reader.Read的辅助性函数。当且仅当err == nil时，返回值n == len(b)
	b := make([]byte, 32)
	m, err := rand.Read(b)
	if err == nil {
		fmt.Println("crypto 随机Read:", b[:m])
		fmt.Println("crypto 随机Read base64:", base64.URLEncoding.EncodeToString(b))
	}
}
