package main

import (
	"fmt"
	"strconv"
)

// 将 10 进制转成 2进制 字符串
func convertToBin(n int) (res string) {
	for ; n > 0; n /= 2 {
		lsb := n % 2
		res = strconv.Itoa(lsb) + res
	}
	return
}
func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(4314),
		convertToBin(0),
	)
}
