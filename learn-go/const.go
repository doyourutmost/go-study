package main

import (
	"fmt"
	"math"
)

// go 语言的不带类型的常量是记录数值的，无关类型
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)
	// b,kb,mb,gb,tb,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, python, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func main() {
	consts()
	enums()
}
