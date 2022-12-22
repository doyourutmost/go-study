package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*bool,string
(u)int,(u)int8,(u)int16,(u)int32,(u)int64,uintptr
byte,rune
float32,float64*/
// 欧拉函数
func eular() {
	fmt.Println(cmplx.Exp(1i*math.Pi)+1, cmplx.Pow(math.E, 1i*math.Pi)+1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

// go 语言没有隐式类型转换，只有强制类型转换
func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

func main() {
	eular()
	triangle()
}
