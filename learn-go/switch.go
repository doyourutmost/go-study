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

func grade(score int) (res string) {
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong Score: %d", score))
	case score < 60:
		res = "F"
	case score < 80:
		res = "C"
	case score < 90:
		res = "B"
	case score <= 100:
		res = "A"
	}
	return
}
func main() {
	fmt.Println(eval(99, 99, '*'),
		eval(88, 33, '/'),
		//eval(77, 33, '=')
	)
	fmt.Println(grade(99),
		grade(88),
		grade(101))

}
