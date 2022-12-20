package main

import "fmt"

// switch会自动break，除非使用 fallthrough
func eval(a, b int, op rune) (res int) {
	switch op {
	case '+':
		res = a + b
	case '-':
		res = a - b
	case '*':
		res = a * b
	case '/':
		res = a / b
	default:
		panic("unsupported operator: " + string(op))
	}
	return
}
func main() {
	fmt.Println(eval(99, 99, '*'),
		eval(88, 33, '/'),
		eval(77, 33, '='))
}
