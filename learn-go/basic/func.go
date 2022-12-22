package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// go 里面函数是一等公民，可以当作参数传递
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d,%d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sumArgs(values ...int) (sum int) {
	for i := range values {
		sum += values[i]
	}
	return
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

// 返回值类型写在最后面
// 可返回多个值
// 函数作为参数
// 没有默认参数，可选参数
func main() {
	fmt.Println(apply(pow, 3, 4),
		apply(func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4),
		apply(func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 4, 3),
	)
	fmt.Println(sumArgs(2, 3, 5))
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
}
